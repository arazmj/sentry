package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/arazmj/sentry-run/api/proto"
	"github.com/arazmj/sentry-run/pkg/jobmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSentryServiceServer
	manager *jobmanager.JobManager
}

func (s *server) StartJob(ctx context.Context, req *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	job, err := s.manager.StartJob(req.Command, req.CommandArgs, req.MemoryLimit, req.CpuLimit, req.Mount, req.WriteBps, req.ReadBps)
	if err != nil {
		log.Printf("Failed to start job %v", err)
		return nil, err
	}
	return &pb.StartJobResponse{
		JobId: job.ID,
	}, nil
}

func (s *server) StopJob(ctx context.Context, req *pb.StopJobRequest) (*pb.StopJobResponse, error) {
	err := s.manager.StopJob(req.JobId)
	if err != nil {
		log.Printf("Failed to stop job %s: %v", req.JobId, err)
		return &pb.StopJobResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	log.Printf("Stopped job %s", req.JobId)
	return &pb.StopJobResponse{
		Success: true,
		Message: fmt.Sprintf("Job %s stopped successfully", req.JobId),
	}, nil
}

func (s *server) GetJobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	isRunning, err := s.manager.GetJobStatus(req.JobId)
	if err != nil {
		log.Printf("Failed to get job status %s: %v", req.JobId, err)
		return &pb.JobStatusResponse{
			IsRunning: false,
		}, nil
	}

	return &pb.JobStatusResponse{
		IsRunning: isRunning,
	}, nil
}

func (s *server) GetJobLogs(ctx context.Context, req *pb.JobLogsRequest) (*pb.JobLogsResponse, error) {
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

func (s *server) ListJobs(ctx context.Context, req *pb.ListJobsRequest) (*pb.ListJobsResponse, error) {
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

func (s *server) KillJob(ctx context.Context, req *pb.KillJobRequest) (*pb.KillJobResponse, error) {
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

func (s *server) StreamJobLogs(req *pb.JobLogsRequest, stream pb.SentryService_StreamJobLogsServer) error {
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

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	log.SetOutput(os.Stdout)

	// Load server certificate and private key
	serverCert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatalf("Failed to load server certificates: %v", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}

	// Create certificate pool and append CA certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		log.Fatal("Failed to append CA certificate to pool")
	}

	// Create TLS credentials
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	creds := credentials.NewTLS(tlsConfig)

	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	srv := &server{
		manager: jobmanager.New(),
	}
	pb.RegisterSentryServiceServer(s, srv)

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start signal handler
	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal. Cleaning up...")
		srv.manager.KillJobsAll()
		os.Exit(0)
	}()

	log.Printf("Server listening on port %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
