# Design Document for Sentry Job Manager
## Overview
Sentry is a gRPC-based job manager that allows users to start, stop, monitor, 
and manage resource-constrained jobs. It provides a CLI interface for interaction 
and enforces resource limits via cgroups.

## Design Approach
The system follows a client-server model with gRPC for communication. 
The server runs jobs within isolated environments using Linux process controls and cgroups, 
while the client provides a CLI for job interaction.

## Scope
1. Start, stop, monitor, and retrieve real-time logs of jobs.
2. Enforce CPU, memory, and disk I/O constraints using cgroups2.
3. Support job execution in a Linux new namespace.
4. Provide a secure gRPC API with TLS authentication.
5. Provide a authorization by a config file representing assigned roles to a user

## Proposed gRPC API

```
StartJob(StartJobRequest) returns (stream JobOutput)

GetJobStatus(JobStatusRequest) returns (JobStatusResponse)

GetJobLogs(JobLogsRequest) returns (JobLogsResponse)

ListJobs(ListJobsRequest) returns (ListJobsResponse)

KillJob(KillJobRequest) returns (KillJobResponse)

StreamJobLogs(JobLogsRequest) returns (stream JobOutput)
```

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
  * **--force**: Stream command output to console
  * **--cpu-limit**: CPU limit in shares in format of QUOTA PERIOD

    QUOTA: The maximum amount of CPU time a cgroup can consume within a scheduling period, in microseconds (µs).

    PERIOD: The length of the scheduling period, also in microseconds (µs).

    Example:
    ```
    50000 100000
    ```

    This means the processes in this cgroup can use 50,000 µs (50ms) of CPU time every 100,000 µs (100ms) period, effectively limiting the CPU usage to 50%.
  * **--memory-limit**: The maximum amount of memory allowed in bytes
  * **--mount**: Directory path to mount for the job. The server does not mount /bin /usr/bin /lib directories automatically 
  * **--rbps-limit**: Read bytes per second limit (e.g., '1048576' for 1MB/s)
  * **--wbps-limit**: Write bytes per second limit (e.g., '1048576' for 1MB/s)


  Example:
  ```
  $ sentry start -cmd "python script.py" -memory-limit "51200000" -cpu-limit "2000000 5000000"
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
  JOB ID     STATUS     MEM-LIMIT  CPU-LIMIT       WRITE-BPS       READ-BPS        COMMAND
  ------------------------------------------------------------------------------------------
  1234       running    30000      10000 20000     120000          120000          dd 
  ```
 
* **logs** Shows the output of the job or streams the running job output

  Parameters:

  **-force**: Stream logs in real-time

  **-id**: Job ID
  
  Example:
  ```
  $ sentry-run logs -id 1234
  Job output logs...
  Error messages...
  ```
  

## Process Execution Lifecycle

### Job Initialization:
* User starts a job via CLI.
* The server authorizes the user against for the request against the roles defined in sentry-security. 
* The server validates request and spawns a new process using `/bin/sh -c`.
* The process is assigned to a cgroup with defined CPU, memory, and I/O constraints.
* Job output is captured and stored in memory and sent to CLI clients.

### Job Execution:
* The process runs within its assigned cgroup.
* Output is streamed to subscribers.
* Status is tracked in memory.

### Job Termination:
* When a stop or kill request is received, the process is terminated via SIGTERM.
* The associated cgroup is cleaned up. If the server shuts down, it ensures that all running jobs are terminated gracefully. A termination signal triggers cleanup procedures that remove jobs from memory, free allocated resources, and delete the corresponding cgroups. If a forced shutdown occurs, any remaining jobs might be left in an inconsistent state, requiring manual cleanup upon restart.
* The job record is removed from the manager.

## Implementation Details
* Server: Implements job control logic using JobManager.
* CLI Client: Sends gRPC requests and handles responses.
* Cgroups Management: Uses /sys/fs/cgroup for process isolation.
* Logging & Streaming: Uses for real-time output streaming.
* Signal Handling: Gracefully handles termination signals and cleans up running jobs.

### User Role Configuration
A `sentry-roles.toml` file defines user roles and allowed requests:
```toml
[users]

[users.alice]
allowed_requests = ["StartJob", "StopJob", "GetJobStatus", "GetJobLogs"]

[users.bob]
allowed_requests = ["StartJob", "ListJobs"]
```

## Edge Cases 
* Ensures proper error handling. 
* Jobs exceeding limits are killed.
* Ensures thread-safe job management.
* Avoid dangling goroutines or memory leaks
* Proper cleanup after job execution finished, cleaning up cgroup file descriptor
* Server handles SIGINT to allow safe termination.


