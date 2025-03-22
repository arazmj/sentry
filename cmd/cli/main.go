package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pb "github.com/arazmj/sentry-run/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [options]")
		fmt.Println("Commands:")
		fmt.Println("  start   Start a new job")
		fmt.Println("  stop    Stop a running job")
		fmt.Println("  status  Get job status")
		fmt.Println("  logs    Get job logs")
		fmt.Println("  list    List all jobs")
		fmt.Println("  kill    Kill a job (SIGKILL)")
		os.Exit(1)
	}

	command := os.Args[1]
	serverAddr := "localhost:50051"

	// Create separate FlagSets for each command
	startFlags := flag.NewFlagSet("start", flag.ExitOnError)
	startCmd := startFlags.String("cmd", "", "Command to execute")

	statusFlags := flag.NewFlagSet("status", flag.ExitOnError)
	statusID := statusFlags.String("id", "", "Job ID")

	logsFlags := flag.NewFlagSet("logs", flag.ExitOnError)
	logsID := logsFlags.String("id", "", "Job ID")
	logsForce := logsFlags.Bool("force", false, "Stream logs in real-time")

	killFlags := flag.NewFlagSet("kill", flag.ExitOnError)
	killID := killFlags.String("id", "", "Job ID")

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start signal handler
	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal. Cleaning up...")
		cancel()
	}()

	// Load client certificate and private key
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatalf("Failed to load client certificates: %v", err)
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
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}
	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient(serverAddr,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSentryServiceClient(conn)

	switch command {
	case "start":
		startFlags.Parse(os.Args[2:])
		if *startCmd == "" {
			log.Fatal("Command is required for start action. Use -cmd flag")
		}

		job, err := client.StartJob(ctx, &pb.StartJobRequest{
			Command:     *startCmd,
			CommandArgs: startFlags.Args(),
		})
		if err != nil {
			log.Fatalf("Could not start job: %v", err)
		}
		fmt.Printf("Job ID: %v\n", job.JobId)
	case "status":
		statusFlags.Parse(os.Args[2:])
		if *statusID == "" {
			log.Fatal("Job ID is required for status action. Use -id flag")
		}
		resp, err := client.GetJobStatus(ctx, &pb.JobStatusRequest{
			JobId: *statusID,
		})
		if err != nil {
			log.Fatalf("Could not get job status: %v", err)
		}

		status := "stopped"
		if resp.IsRunning {
			status = "running"
		}
		fmt.Printf("Job status: %s\n", status)

	case "logs":
		logsFlags.Parse(os.Args[2:])
		if *logsID == "" {
			log.Fatal("Job ID is required for logs action. Use -id flag")
		}

		stream, err := client.StreamJobLogs(ctx, &pb.JobLogsRequest{
			JobId: *logsID,
		})
		if err != nil {
			log.Fatalf("Could not stream logs: %v", err)
		}

		for {
			select {
			case <-ctx.Done():
			default:
				output, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					if errors.Is(ctx.Err(), context.Canceled) {
						return
					}
					log.Printf("Error receiving logs: %v", err)
					return
				}

				if output.IsStderr {
					_, _ = fmt.Fprintf(os.Stderr, "%s", output.Data)
				} else {
					fmt.Printf("%s", output.Data)
				}
			}
			if !*logsForce {
				return
			}
		}
	case "list":
		resp, err := client.ListJobs(ctx, &pb.ListJobsRequest{})
		if err != nil {
			log.Fatalf("Could not list jobs: %v", err)
		}

		if len(resp.Jobs) == 0 {
			fmt.Println("No jobs found")
			return
		}

		format := "%-40s %-10s %s\n"
		fmt.Printf(format, "JOB ID", "STATUS", "COMMAND")
		fmt.Println(strings.Repeat("-", 60))

		for _, job := range resp.Jobs {
			status := "stopped"
			if job.IsRunning {
				status = "running"
			}

			fmt.Printf(format, job.JobId, status, job.Command)
		}

	case "kill":
		killFlags.Parse(os.Args[2:])
		if *killID == "" {
			log.Fatal("Job ID is required for kill action. Use -id flag")
		}
		resp, err := client.KillJob(ctx, &pb.KillJobRequest{
			JobId: *killID,
		})
		if err != nil {
			log.Fatalf("Could not kill job: %v", err)
		}
		fmt.Printf("Kill job result: %s\n", resp.Message)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
