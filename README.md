# Sentry Run

A secure job management system that allows running and monitoring processes with resource constraints using cgroups on Linux systems. The system uses gRPC with mutual TLS authentication for secure client-server communication.

## Features

- Start jobs with resource constraints:
    - Memory limits
    - CPU limits
    - I/O bandwidth limits (read/write BPS)
    - Directory mounting (chroot)
- Real-time job monitoring
- Secure communication using mutual TLS
- Job management operations:
    - Start jobs
    - Stop jobs gracefully
    - Kill jobs (SIGKILL)
    - Get job status
    - Stream job logs in real-time
    - List all running jobs

## Prerequisites

- Linux system with cgroups v2
- Go 1.x
- Valid TLS certificates (client, server, and CA certificates)

## Installation

```bash
go get github.com/arazmj/sentry-run
```

## Certificate Setup

Place your TLS certificates in the `certs` directory:
- `certs/server.crt` and `certs/server.key` - Server certificate and private key
- `certs/client.crt` and `certs/client.key` - Client certificate and private key
- `certs/ca.crt` - Certificate Authority (CA) certificate

## Usage

### Starting the Server

```bash
./server
```

The server listens on port 50051 by default.

### Using the CLI

The CLI supports the following commands:

```bash
# Start a new job
cli start -cmd "your_command" [options]
Options:
  -memory-limit string   Memory limit (e.g., '100M', '1G')
  -cpu-limit string     CPU limit in shares (e.g., '512')
  -mount string         Directory path to mount for the job
  -wbps-limit string    Write bytes per second limit (e.g., '1048576' for 1MB/s)
  -rbps-limit string    Read bytes per second limit (e.g., '1048576' for 1MB/s)

# Get job status
cli status -id <job_id>

# Stream job logs
cli logs -id <job_id> [-force]
Options:
  -force    Stream logs in real-time

# List all jobs
cli list

# Kill a job
cli kill -id <job_id>
```

## Example

```bash
# Start a memory-limited job
cli start -cmd "stress --vm 1 --vm-bytes 50M" -memory-limit 100M

# Monitor job logs in real-time
cli logs -id <job_id> -force

# List all running jobs
cli list
```

## Security

The system uses mutual TLS authentication to ensure secure communication between the client and server. Both the client and server must present valid certificates signed by the trusted CA.
