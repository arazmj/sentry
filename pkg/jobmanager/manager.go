package jobmanager

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
)

type Job struct {
	ID            string
	PID           int
	Command       string
	Cmd           *exec.Cmd
	Stdout        io.ReadCloser
	Stderr        io.ReadCloser
	stdoutHistory bytes.Buffer
	stderrHistory bytes.Buffer
	callbacks     []OutputCallback
	MemoryLimit   string
	CpuLimit      string
	Mount         string
	WriteBps      string
	ReadBps       string
	DeviceId      string
}

type JobManager struct {
	jobs sync.Map
}

// OutputCallback is called for each line of output from the job
type OutputCallback func(data []byte, isStderr bool) error

// New creates a new JobManager
func New() *JobManager {
	return &JobManager{}
}

const (
	cgroupBasePath = "/sys/fs/cgroup"
	cgroupName     = "sentry-run"
)

func getCgroupPath(jobID string) string {
	return filepath.Join(cgroupBasePath, fmt.Sprintf("%s-%s", cgroupName, jobID))
}

func setLimits(pid int, jobID, cpuLimit, memoryLimit, writeBps, readBps string) error {
	cgroupPath := getCgroupPath(jobID)

	if err := os.MkdirAll(cgroupPath, 0755); err != nil {
		return fmt.Errorf("failed to create cgroup directory %s: %v", cgroupPath, err)
	}

	// Associate process with cgroup
	procsPath := filepath.Join(cgroupPath, "cgroup.procs")
	if err := os.WriteFile(procsPath, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("failed to add process to CPU cgroup: %v", err)

	}

	if cpuLimit != "" {
		cpuPath := filepath.Join(cgroupPath, "cpu.max")
		if err := os.WriteFile(cpuPath, []byte(cpuLimit), 0644); err != nil {
			return fmt.Errorf("failed to set CPU cpu limit: %v", err)
		}
	}

	if memoryLimit != "" {
		memoryPath := filepath.Join(cgroupPath, "memory.max")
		if err := os.WriteFile(memoryPath, []byte(memoryLimit), 0644); err != nil {
			return fmt.Errorf("failed to set memory memoryLimit: %v", err)
		}
	}

	if writeBps != "" || readBps != "" {
		// Get file system stats for the current working directory
		var stat syscall.Stat_t
		err := syscall.Stat("/", &stat) // Root filesystem
		if err != nil {
			log.Fatalf("Failed to stat filesystem: %v", err)
		}

		// Extract major and minor device numbers
		major := (stat.Dev >> 8) & 0xfff // Major number
		minor := 0                       // Minor should be set to zero because cgroups only works on whole device block

		ioLimit := fmt.Sprintf("%d:%d", major, minor)

		ioPath := filepath.Join(cgroupPath, "io.max")
		if writeBps != "" {
			ioLimit += fmt.Sprintf(" wbps=%s", writeBps)
		}
		if readBps != "" {
			ioLimit += fmt.Sprintf(" rbps=%s", readBps)
		}

		if err := os.WriteFile(ioPath, []byte(ioLimit), 0644); err != nil {
			return fmt.Errorf("failed to set io limit: %v", err)
		}
	}
	return nil
}

func cleanupCgroup(job *Job) error {
	cgroupPath := getCgroupPath(job.ID)

	// The cgroup dir can only be deleted by rmdir syscall
	if err := syscall.Rmdir(cgroupPath); err != nil {
		return fmt.Errorf("failed to remove cgroup directory %s: %v", cgroupPath, err)
	}

	return nil
}

// StartJob starts a new job and returns its ID
func (m *JobManager) StartJob(command string, commandArgs []string, memoryLimit, cpuLimit, mount, writeBps, readBps string) (*Job, error) {
	cmd := exec.Command(command, commandArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if mount != "" {
		//cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWNS
		cmd.SysProcAttr.Chroot = mount
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		stdout.Close()
		return nil, fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		stdout.Close()
		stderr.Close()
		return nil, fmt.Errorf("failed to start command: %v", err)
	}

	jobUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate job uuid: %v", err)
		stdout.Close()
		stderr.Close()
	}
	jobID := jobUUID.String()

	if err := setLimits(cmd.Process.Pid, jobID, cpuLimit, memoryLimit, writeBps, readBps); err != nil {
		err := cmd.Process.Kill()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to set limits: %v", err)
	}

	job := &Job{
		ID:          jobID,
		PID:         cmd.Process.Pid,
		Command:     command,
		Cmd:         cmd,
		Stdout:      stdout,
		Stderr:      stderr,
		callbacks:   []OutputCallback{},
		MemoryLimit: memoryLimit,
		CpuLimit:    cpuLimit,
		Mount:       mount,
		WriteBps:    writeBps,
		ReadBps:     readBps,
	}

	go job.serveSubscribers()

	m.jobs.Store(job.ID, job)
	return job, nil
}

// StopJob stops a running job
func (m *JobManager) StopJob(jobID string) error {
	value, exists := m.jobs.LoadAndDelete(jobID)
	if !exists {
		return fmt.Errorf("job %s not found", jobID)
	}

	job := value.(*Job)
	if err := job.Cmd.Process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to stop job %s: %v", jobID, err)
	}

	job.Stdout.Close()
	job.Stderr.Close()

	// Clean up cgroups
	if err := cleanupCgroup(job); err != nil {
		log.Printf("Warning: failed to clean up cgroups for job %s: %v", jobID, err)
	}

	return nil
}

// GetJobStatus returns whether the job is running
func (m *JobManager) GetJobStatus(jobID string) (bool, error) {
	value, exists := m.jobs.Load(jobID)
	if !exists {
		return false, fmt.Errorf("job %s not found", jobID)
	}

	job := value.(*Job)
	if err := job.Cmd.Process.Signal(syscall.Signal(0)); err != nil {
		return false, nil
	}

	return true, nil
}

func (m *JobManager) StreamOutput(ctx context.Context, jobID string, callback OutputCallback) error {
	value, exists := m.jobs.Load(jobID)
	if !exists {
		return fmt.Errorf("job %s not found", jobID)
	}

	job := value.(*Job)

	// Add callback for future output
	job.callbacks = append(job.callbacks, callback)

	doneCh := make(chan error, 1)

	go func() {
		doneCh <- job.Cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		// context canceled, return immediately
		return ctx.Err()
	case err := <-doneCh:
		// process exited naturally
		return err
	}
}

func (job *Job) serveSubscribers() {
	var wg sync.WaitGroup
	wg.Add(2)

	streamOutput := func(reader io.Reader, history *bytes.Buffer, isStderr bool) {
		defer wg.Done()
		buffer := make([]byte, 32*1024)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				data := buffer[:n]
				history.Write(data)
				for _, callback := range job.callbacks {
					err := callback(data, isStderr)
					if err != nil {
						log.Printf("Warning: failed to send callback: %v", err)
						return
					}
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error reading stdout: %v", err)
				break
			}
		}
	}

	go streamOutput(job.Stdout, &job.stdoutHistory, false)
	go streamOutput(job.Stderr, &job.stderrHistory, true)

	// Wait for both streams to complete in a separate goroutine
	wg.Wait()

	// Clean up after the job is complete
	err := cleanupCgroup(job)
	if err != nil {
		log.Printf("Warning: failed to cleanup cgroup for job %s: %v", job.ID, err)
	}

	job.Stdout.Close()
	job.Stderr.Close()
}

// GetJobOutput returns the output history of a job
func (m *JobManager) GetJobOutput(jobID string) (stdout, stderr []byte, err error) {
	value, exists := m.jobs.Load(jobID)
	if !exists {
		return nil, nil, fmt.Errorf("job %s not found", jobID)
	}

	job := value.(*Job)
	return job.stdoutHistory.Bytes(), job.stderrHistory.Bytes(), nil
}

// ListJobs returns a list of all jobs
func (m *JobManager) ListJobs() []*Job {
	var jobs []*Job
	m.jobs.Range(func(key, value interface{}) bool {
		job := value.(*Job)
		jobs = append(jobs, job)
		return true
	})
	return jobs
}

func (m *JobManager) KillJob(jobID string) error {
	value, exists := m.jobs.LoadAndDelete(jobID)
	if !exists {
		return fmt.Errorf("job %s not found", jobID)
	}

	job := value.(*Job)

	if err := job.Cmd.Process.Signal(syscall.SIGKILL); err != nil {
		return fmt.Errorf("failed to kill job %s: %v", jobID, err)
	}

	err := cleanupCgroup(job)
	if err != nil {
		return err
	}

	job.Stdout.Close()
	job.Stderr.Close()

	return nil
}

func (m *JobManager) KillJobsAll() {
	m.jobs.Range(func(key, value interface{}) bool {
		job := value.(*Job)
		m.KillJob(job.ID)
		return true
	})
}
