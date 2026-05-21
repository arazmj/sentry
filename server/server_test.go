package main

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	pb "github.com/arazmj/sentry-run/api/proto"
	"github.com/arazmj/sentry-run/pkg/jobmanager"
)

type fakeJobManager struct {
	startJob    *jobmanager.Job
	startErr    error
	startCall   startJobCall
	stopErr     error
	stoppedJob  string
	status      map[string]bool
	statusErr   error
	statusCalls []string
	stdout      []byte
	stderr      []byte
	outputJob   string
	jobs        []*jobmanager.Job
	killErr     error
	killedJob   string
}

type startJobCall struct {
	command     string
	commandArgs []string
	memoryLimit string
	cpuLimit    string
	mount       string
	writeBps    string
	readBps     string
}

func (f *fakeJobManager) StartJob(command string, commandArgs []string, memoryLimit, cpuLimit, mount, writeBps, readBps string) (*jobmanager.Job, error) {
	f.startCall = startJobCall{command, commandArgs, memoryLimit, cpuLimit, mount, writeBps, readBps}
	return f.startJob, f.startErr
}
func (f *fakeJobManager) StopJob(jobID string) error { f.stoppedJob = jobID; return f.stopErr }
func (f *fakeJobManager) KillJob(jobID string) error { f.killedJob = jobID; return f.killErr }
func (f *fakeJobManager) GetJobStatus(jobID string) (bool, error) {
	f.statusCalls = append(f.statusCalls, jobID)
	if f.statusErr != nil {
		return false, f.statusErr
	}
	return f.status[jobID], nil
}
func (f *fakeJobManager) GetJobOutput(jobID string) ([]byte, []byte, error) {
	f.outputJob = jobID
	return f.stdout, f.stderr, nil
}
func (f *fakeJobManager) ListJobs() []*jobmanager.Job { return f.jobs }
func (f *fakeJobManager) StreamOutput(ctx context.Context, jobID string, callback jobmanager.OutputCallback) error {
	return nil
}

func TestStartJobSuccess(t *testing.T) {
	fake := &fakeJobManager{startJob: &jobmanager.Job{ID: "job-1"}}
	req := &pb.StartJobRequest{Command: "echo", CommandArgs: []string{"hello"}, MemoryLimit: "128M", CpuLimit: "10000 100000", Mount: "/srv", WriteBps: "1M", ReadBps: "2M"}
	resp, err := NewServer(fake).StartJob(context.Background(), req)
	if err != nil {
		t.Fatalf("StartJob returned error: %v", err)
	}
	if resp.GetJobId() != "job-1" {
		t.Fatalf("JobId = %q, want job-1", resp.GetJobId())
	}
	want := startJobCall{req.Command, req.CommandArgs, req.MemoryLimit, req.CpuLimit, req.Mount, req.WriteBps, req.ReadBps}
	if !reflect.DeepEqual(fake.startCall, want) {
		t.Fatalf("StartJob call = %#v, want %#v", fake.startCall, want)
	}
}

func TestStartJobError(t *testing.T) {
	boom := errors.New("boom")
	resp, err := NewServer(&fakeJobManager{startErr: boom}).StartJob(context.Background(), &pb.StartJobRequest{Command: "false"})
	if !errors.Is(err, boom) {
		t.Fatalf("error = %v, want %v", err, boom)
	}
	if resp != nil {
		t.Fatalf("response = %#v, want nil", resp)
	}
}

func TestStopJobSuccess(t *testing.T) {
	fake := &fakeJobManager{}
	resp, err := NewServer(fake).StopJob(context.Background(), &pb.StopJobRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("StopJob returned error: %v", err)
	}
	if !resp.GetSuccess() || !strings.Contains(resp.GetMessage(), "job-1") {
		t.Fatalf("response = %#v, want success", resp)
	}
	if fake.stoppedJob != "job-1" {
		t.Fatalf("StopJob called with %q", fake.stoppedJob)
	}
}

func TestStopJobError(t *testing.T) {
	boom := errors.New("stop failed")
	resp, err := NewServer(&fakeJobManager{stopErr: boom}).StopJob(context.Background(), &pb.StopJobRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("handler error = %v", err)
	}
	if resp.GetSuccess() || resp.GetMessage() != boom.Error() {
		t.Fatalf("response = %#v, want failure", resp)
	}
}

func TestGetJobStatusSuccess(t *testing.T) {
	resp, err := NewServer(&fakeJobManager{status: map[string]bool{"job-1": true}}).GetJobStatus(context.Background(), &pb.JobStatusRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("GetJobStatus returned error: %v", err)
	}
	if !resp.GetIsRunning() {
		t.Fatal("IsRunning = false, want true")
	}
}

func TestGetJobStatusError(t *testing.T) {
	resp, err := NewServer(&fakeJobManager{statusErr: errors.New("missing")}).GetJobStatus(context.Background(), &pb.JobStatusRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("handler error = %v", err)
	}
	if resp.GetIsRunning() {
		t.Fatal("IsRunning = true, want false")
	}
}

func TestGetJobLogsSuccess(t *testing.T) {
	fake := &fakeJobManager{stdout: []byte("out"), stderr: []byte("err")}
	resp, err := NewServer(fake).GetJobLogs(context.Background(), &pb.JobLogsRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("GetJobLogs returned error: %v", err)
	}
	want := "=== STDOUT ===\nout\n=== STDERR ===\nerr"
	if string(resp.GetLogs()) != want {
		t.Fatalf("logs = %q, want %q", string(resp.GetLogs()), want)
	}
	if fake.outputJob != "job-1" {
		t.Fatalf("GetJobOutput called with %q", fake.outputJob)
	}
}

func TestListJobsEmpty(t *testing.T) {
	resp, err := NewServer(&fakeJobManager{}).ListJobs(context.Background(), &pb.ListJobsRequest{})
	if err != nil {
		t.Fatalf("ListJobs returned error: %v", err)
	}
	if len(resp.GetJobs()) != 0 {
		t.Fatalf("got %d jobs, want 0", len(resp.GetJobs()))
	}
}

func TestListJobsMultiple(t *testing.T) {
	fake := &fakeJobManager{
		jobs: []*jobmanager.Job{
			{ID: "job-1", Command: "echo", MemoryLimit: "128M", CpuLimit: "10000 100000", Mount: "/srv", ReadBps: "1M", WriteBps: "2M"},
			{ID: "job-2", Command: "sleep"},
		},
		status: map[string]bool{"job-1": true, "job-2": false},
	}
	resp, err := NewServer(fake).ListJobs(context.Background(), &pb.ListJobsRequest{})
	if err != nil {
		t.Fatalf("ListJobs returned error: %v", err)
	}
	if len(resp.GetJobs()) != 2 {
		t.Fatalf("got %d jobs, want 2", len(resp.GetJobs()))
	}
	if got := resp.GetJobs()[0]; got.GetJobId() != "job-1" || got.GetCommand() != "echo" || !got.GetIsRunning() || got.GetMemoryLimit() != "128M" || got.GetCpuLimit() != "10000 100000" || got.GetMount() != "/srv" || got.GetReadBps() != "1M" || got.GetWriteBps() != "2M" {
		t.Fatalf("first job = %#v", got)
	}
	if got := resp.GetJobs()[1]; got.GetJobId() != "job-2" || got.GetCommand() != "sleep" || got.GetIsRunning() {
		t.Fatalf("second job = %#v", got)
	}
	if !reflect.DeepEqual(fake.statusCalls, []string{"job-1", "job-2"}) {
		t.Fatalf("status calls = %v", fake.statusCalls)
	}
}

func TestKillJobSuccess(t *testing.T) {
	fake := &fakeJobManager{}
	resp, err := NewServer(fake).KillJob(context.Background(), &pb.KillJobRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("KillJob returned error: %v", err)
	}
	if !resp.GetSuccess() || !strings.Contains(resp.GetMessage(), "job-1") {
		t.Fatalf("response = %#v, want success", resp)
	}
	if fake.killedJob != "job-1" {
		t.Fatalf("KillJob called with %q", fake.killedJob)
	}
}

func TestKillJobError(t *testing.T) {
	boom := errors.New("kill failed")
	resp, err := NewServer(&fakeJobManager{killErr: boom}).KillJob(context.Background(), &pb.KillJobRequest{JobId: "job-1"})
	if err != nil {
		t.Fatalf("handler error = %v", err)
	}
	if resp.GetSuccess() || resp.GetMessage() != boom.Error() {
		t.Fatalf("response = %#v, want failure", resp)
	}
}
