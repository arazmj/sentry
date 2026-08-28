//go:build linux || darwin

package jobmanager

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestNewReturnsManager(t *testing.T) {
	if New() == nil {
		t.Fatal("New() returned nil")
	}
}

func TestStartJobNoLimitsCapturesOutputAndListsJob(t *testing.T) {
	manager := New()
	job, err := manager.StartJob("/bin/echo", []string{"hello"}, "", "", "", "", "")
	if err != nil {
		t.Fatalf("StartJob() error = %v", err)
	}
	defer func() {
		if err := job.Cmd.Wait(); err != nil {
			t.Logf("waiting for job exit: %v", err)
		}
	}()

	if job.ID == "" {
		t.Fatal("StartJob() returned job with empty ID")
	}
	if job.Command != "/bin/echo" {
		t.Fatalf("job.Command = %q, want /bin/echo", job.Command)
	}
	if job.PID <= 0 {
		t.Fatalf("job.PID = %d, want > 0", job.PID)
	}

	var stdout []byte
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		stdout, _, err = manager.GetJobOutput(job.ID)
		if err != nil {
			t.Fatalf("GetJobOutput() error = %v", err)
		}
		if strings.Contains(string(stdout), "hello") {
			break
		}
		time.Sleep(25 * time.Millisecond)
	}
	if !strings.Contains(string(stdout), "hello") {
		t.Fatalf("stdout = %q, want to contain hello", string(stdout))
	}

	jobs := manager.ListJobs()
	for _, listed := range jobs {
		if listed.ID == job.ID {
			return
		}
	}
	t.Fatalf("ListJobs() did not include job %s", job.ID)
}

func TestUnknownJobErrors(t *testing.T) {
	manager := New()
	unknownID := "missing-job"

	if running, err := manager.GetJobStatus(unknownID); err == nil || running {
		t.Fatalf("GetJobStatus(%q) = (%v, %v), want error and running=false", unknownID, running, err)
	}
	if err := manager.StopJob(unknownID); err == nil {
		t.Fatalf("StopJob(%q) succeeded, want error", unknownID)
	}
	if err := manager.KillJob(unknownID); err == nil {
		t.Fatalf("KillJob(%q) succeeded, want error", unknownID)
	}
	if stdout, stderr, err := manager.GetJobOutput(unknownID); err == nil || stdout != nil || stderr != nil {
		t.Fatalf("GetJobOutput(%q) = (%q, %q, %v), want nil output and error", unknownID, stdout, stderr, err)
	}
	if err := manager.StreamOutput(context.Background(), unknownID, func([]byte, bool) error { return nil }); err == nil {
		t.Fatalf("StreamOutput(%q) succeeded, want error", unknownID)
	}
}

func TestGetJobStatusReturnsFalseAfterExit(t *testing.T) {
	manager := New()
	job, err := manager.StartJob("/bin/sh", []string{"-c", "exit 0"}, "", "", "", "", "")
	if err != nil {
		t.Fatalf("StartJob() error = %v", err)
	}
	if err := job.Cmd.Wait(); err != nil {
		t.Fatalf("waiting for job exit: %v", err)
	}

	running, err := manager.GetJobStatus(job.ID)
	if err != nil {
		t.Fatalf("GetJobStatus() error = %v", err)
	}
	if running {
		t.Fatal("GetJobStatus() running = true, want false after process exits")
	}
}
