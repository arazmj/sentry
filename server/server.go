package main

import (
	"context"
	"fmt"
	"log/slog"

	pb "github.com/arazmj/sentry-run/api/proto"
	"github.com/arazmj/sentry-run/pkg/jobmanager"
)

type JobManager interface {
	StartJob(command string, commandArgs []string, memoryLimit, cpuLimit, mount, writeBps, readBps string) (*jobmanager.Job, error)
	StopJob(jobID string) error
	KillJob(jobID string) error
	GetJobStatus(jobID string) (bool, error)
	GetJobOutput(jobID string) (stdout, stderr []byte, err error)
	ListJobs() []*jobmanager.Job
	StreamOutput(ctx context.Context, jobID string, callback jobmanager.OutputCallback) error
}

type Server struct {
	pb.UnimplementedSentryServiceServer
	manager JobManager
}

func NewServer(manager JobManager) *Server {
	return &Server{manager: manager}
}

func (s *Server) StartJob(ctx context.Context, req *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	job, err := s.manager.StartJob(req.Command, req.CommandArgs, req.MemoryLimit, req.CpuLimit, req.Mount, req.WriteBps, req.ReadBps)
	if err != nil {
		slog.Error("failed to start job", "error", err)
		return nil, err
	}
	return &pb.StartJobResponse{
		JobId: job.ID,
	}, nil
}

func (s *Server) StopJob(ctx context.Context, req *pb.StopJobRequest) (*pb.StopJobResponse, error) {
	err := s.manager.StopJob(req.JobId)
	if err != nil {
		slog.Error("failed to stop job", "job_id", req.JobId, "error", err)
		return &pb.StopJobResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	slog.Info("stopped job", "job_id", req.JobId)
	return &pb.StopJobResponse{
		Success: true,
		Message: fmt.Sprintf("Job %s stopped successfully", req.JobId),
	}, nil
}

func (s *Server) GetJobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	isRunning, err := s.manager.GetJobStatus(req.JobId)
	if err != nil {
		slog.Error("failed to get job status", "job_id", req.JobId, "error", err)
		return &pb.JobStatusResponse{
			IsRunning: false,
		}, nil
	}

	return &pb.JobStatusResponse{
		IsRunning: isRunning,
	}, nil
}

func (s *Server) GetJobLogs(ctx context.Context, req *pb.JobLogsRequest) (*pb.JobLogsResponse, error) {
	stdout, stderr, err := s.manager.GetJobOutput(req.JobId)
	if err != nil {
		return nil, err
	}

	// Combine stdout and stderr with markers
	logs := fmt.Sprintf("=== STDOUT ===\n%s\n=== STDERR ===\n%s", stdout, stderr)
	return &pb.JobLogsResponse{
		Logs: []byte(logs),
	}, nil
}

func (s *Server) ListJobs(ctx context.Context, req *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
	jobs := s.manager.ListJobs()
	response := &pb.ListJobsResponse{
		Jobs: make([]*pb.JobInfo, 0, len(jobs)),
	}

	for _, job := range jobs {
		isRunning, _ := s.manager.GetJobStatus(job.ID)
		jobInfo := &pb.JobInfo{
			JobId:       job.ID,
			Command:     job.Command,
			IsRunning:   isRunning,
			MemoryLimit: job.MemoryLimit,
			CpuLimit:    job.CpuLimit,
			Mount:       job.Mount,
			ReadBps:     job.ReadBps,
			WriteBps:    job.WriteBps,
		}
		response.Jobs = append(response.Jobs, jobInfo)
	}

	return response, nil
}

func (s *Server) KillJob(ctx context.Context, req *pb.KillJobRequest) (*pb.KillJobResponse, error) {
	err := s.manager.KillJob(req.JobId)
	if err != nil {
		return &pb.KillJobResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.KillJobResponse{
		Success: true,
		Message: fmt.Sprintf("Job %s killed successfully", req.JobId),
	}, nil
}

func (s *Server) StreamJobLogs(req *pb.JobLogsRequest, stream pb.SentryService_StreamJobLogsServer) error {
	// First, send existing logs
	stdout, stderr, err := s.manager.GetJobOutput(req.JobId)
	if err != nil {
		return err
	}

	// Send stdout history
	if len(stdout) > 0 {
		if err := stream.Send(&pb.JobOutput{
			JobId:    req.JobId,
			Data:     stdout,
			IsStderr: false,
		}); err != nil {
			return err
		}
	}

	// Send stderr history
	if len(stderr) > 0 {
		if err := stream.Send(&pb.JobOutput{
			JobId:    req.JobId,
			Data:     stderr,
			IsStderr: true,
		}); err != nil {
			return err
		}
	}

	ctx := stream.Context()

	// Then start streaming new output
	return s.manager.StreamOutput(ctx, req.JobId, func(data []byte, isStderr bool) error {
		return stream.Send(&pb.JobOutput{
			JobId:    req.JobId,
			Data:     data,
			IsStderr: isStderr,
		})
	})
}
