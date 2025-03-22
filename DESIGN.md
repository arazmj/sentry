# Design Document for Sentry Job Manager
## Overview
Sentry is a gRPC-based job manager that allows users to start, stop, monitor, 
and manage resource-constrained jobs. It provides a CLI interface for interaction 
and enforces resource limits via cgroups.

## Scope
* Start, stop, monitor, and retrieve real-time logs of jobs.
* Enforce CPU, memory, and disk I/O constraints using cgroups2 (the limits are hard-coded).
* Support job execution in a Linux new namespace.
* Provide a secure gRPC API with TLS authentication.
* Provide an authorization by a config file representing assigned roles to a user.
* The code is only tested on Debian 6.1.129

## Security Considerations
* The service implements mTLS authentication, requiring both server and client certificates for all connections.
  * Protocol Version: TLS 1.3 
  * Cipher Suites: Using TLS 1.3's default AEAD ciphers:
    * TLS_AES_128_GCM_SHA256
    * TLS_AES_256_GCM_SHA384
    * TLS_CHACHA20_POLY1305_SHA256
  * Elliptic Curves: Default Go preferences:
    * X25519 (primary)
    * NIST P-256
    * NIST P-384
    * NIST P-521
  * RSA-4096 or ECDSA P-256/P-384 keys
  * SHA-256 or stronger signatures
* Only authorized clients with valid certificates can interact with the server.
* Only clients with proper access role can perform an action, i.e Start, Kill

## Proposed gRPC API

```
service SentryService {
  rpc StartJob (StartJobRequest) returns (StartJobResponse)  {}
  rpc KillJob (KillJobRequest) returns (KillJobResponse) {}
  rpc GetJobStatus (JobStatusRequest) returns (JobStatusResponse) {}
  rpc StreamJobLogs (JobLogsRequest) returns (stream JobOutput) {}
  rpc ListJobs (ListJobsRequest) returns (ListJobsResponse) {}
}

message StartJobRequest {
  string command = 1;
  repeated string commandArgs = 2;
}

message JobOutput {
  string job_id = 1;
  bytes data = 2;
  bool is_stderr = 3;
}

message StartJobResponse{
  string job_id = 1;
}

message StopJobRequest {
  string job_id = 1;
}

message StopJobResponse {
  bool success = 1;
  string message = 2;
}

message JobStatusRequest {
  string job_id = 1;
}

message JobStatusResponse {
  bool is_running = 1;
}

message JobLogsRequest {
  string job_id = 1;
}

message JobLogsResponse {
  bytes logs = 1;
}

message ListJobsRequest {}

message JobInfo {
  string job_id = 1;
  string command = 2;
  bool is_running = 3;
}

message ListJobsResponse {
  repeated JobInfo jobs = 1;
}

message KillJobRequest {
  string job_id = 1;
}

message KillJobResponse {
  bool success = 1;
  string message = 2;
}
```

## Job IDs
Jobs are identified are tracked internally as UUIDs rather than process IDs or sequential counters.

## Constant
* **CPU_LIMIT**: CPU limit in shares in format of QUOTA PERIOD

  QUOTA: The maximum amount of CPU time a cgroup can consume within a scheduling period, in microseconds (µs).

  PERIOD: The length of the scheduling period, also in microseconds (µs).

  Example:
   ```
   50000 100000
   ```

  This means the processes in this cgroup can use 50,000 µs (50ms) of CPU time every 100,000 µs (100ms) period, effectively limiting the CPU usage to 50%.
* **MEMORY-LIMIT**: The maximum amount of memory allowed in bytes
* **MOUNT**: Directory path to mount for the job. The server does not mount /bin /usr/bin /lib directories automatically
* **RBPS-LIMIT**: Read bytes per second limit (e.g., '1048576' for 1MB/s)
* **WBPS-LIMIT**: Write bytes per second limit (e.g., '1048576' for 1MB/s)


## CLI User Experience
Users interact via a CLI tool with the following commands:
```
    Usage: cli <command> [options]
    Commands:
      start   Start a new job
      status  Get job status
      logs    Get job logs
      list    List all jobs
      kill    Kill a job (SIGKILL)
```

* **start**: Runs a new job

Usage of start:
 
  Example:
  ```
  $ sentry start -cmd python -- script.py
  Started job with ID: 1234
  
  $ sentry status -id 1234
  Job 1234 is running
  ```

* **status**: Shows the status of the job

  Parameters: -id Job ID

* **kill**: Terminates the job by SIGTERM signal

  Parameters: -id Job ID

* **list**: Lists all running jobs along with assigned parameters 

  Example:
  ```
  JOB ID                                   STATUS     COMMAND
  ------------------------------------------------------------
  6aeab32e-0742-11f0-b330-8e0b0e62b7b5     running    bash
  ```
 
* **logs** Shows the output of the job or streams the running job output

  Parameters:

  **-force**: Stream logs in real-time

  **-id**: Job ID
  
  Example:
  ```
  $ sentry-run logs -id 6aeab32e-0742-11f0-b330-8e0b0e62b7b5
  Job output logs...
  Error messages...
  ```
  

## Process Execution Lifecycle

### Job Initialization:
* User starts a job via CLI.
* The server authorizes the user against for the request against the roles defined in sentry-security. 
  * The client identity is the CN field of the certificate
    * For testing purposes, both the client and server are using a self-signed CA. However, in production, the client should present a user certificate signed by a trusted CA, not the CA itself.
  * The action will run if the client identity is defined in `SENTRY_ROLES` static variable and the user has permission for the action.
* The process is assigned to a cgroup with defined CPU, memory, and I/O constraints.
* Job output is captured and stored in memory and sent to CLI clients.
  * Once the job started a running goroutine in the background reads the stdout/stdout and calls `stream.Send()` for each client them to clients.

### Job Execution:
* The process runs within its assigned cgroup by following process.
  1. The server creates a new directory under /sys/fs/cgroup/sentry-[JobID].
  2. The CPU limit parameter value is written to `cpu.max` fd.
  3. The memory limit parameter value is written to `memory.max` fd.
  4. The disk IO limit parameter value is written to `io.max` fd.
  5. The PID of the job is written to `cgroup.procs` fd.
  
* Output is streamed to subscribers. 
  * The server keeps streaming to all running clients and stops when there is a network transportation error (client disconnects)
  * The system uses a callback-based streaming model to handle process output.
    * Each job maintains a list of output callbacks for real-time streaming
    * Process stdout/stderr are read in separate goroutines using buffered I/O
    * Output is stored in memory buffers for history while simultaneously streaming to active subscribers
* Status is tracked in memory.

### Job Termination:
* To ensure all processes in a spawned group are terminated, we:
  * Assign the process to a new process group using `syscall.SysProcAttr{Setpgid: true}`.
  * Retrieve the process group ID (PGID) using `syscall.Getpgid()`.
  * Send the SIGKILL signal to the process group (-PGID).

* The associated cgroup is cleaned up. If the server shuts down, it ensures that all running jobs are terminated gracefully. A termination signal triggers cleanup procedures that remove jobs from memory, free allocated resources, and delete the corresponding cgroups. If a forced shutdown occurs, any remaining jobs might be left in an inconsistent state, requiring manual cleanup upon restart.
* The job record is removed from the manager.

## Implementation Details
* Server: Implements job control logic using JobManager.
* CLI Client: Sends gRPC requests and handles responses.
* Cgroups Management: Uses `/sys/fs/cgroup` for process isolation.
* Logging & Streaming: Uses for real-time output streaming.
* Signal Handling: Gracefully handles termination signals and cleans up running jobs.

### User Role Configuration
A SENTRY_ROLES const defines user roles and allowed requests:
```toml
alice = ["StartJob", "StopJob", "GetJobStatus", "GetJobLogs"]
bob = ["StartJob", "ListJobs"]
```

## Edge Cases 
* Ensures proper error handling. 
* Jobs exceeding limits are killed.
* Ensures thread-safe job management.
* Avoid dangling goroutines or memory leaks
* Proper cleanup after job execution finished, cleaning up cgroup file descriptor
* Server handles SIGINT to allow safe termination.