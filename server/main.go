package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	pb "github.com/arazmj/sentry-run/api/proto"
	"github.com/arazmj/sentry-run/pkg/jobmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	log.SetOutput(os.Stdout)

	const (
		defaultPort       = 50051
		defaultServerCert = "certs/server.crt"
		defaultServerKey  = "certs/server.key"
		defaultCACert     = "certs/ca.crt"
	)

	port := flag.Int("port", defaultPort, "Port to listen on")
	serverCertPath := flag.String("cert", defaultServerCert, "Server certificate path")
	serverKeyPath := flag.String("key", defaultServerKey, "Server private key path")
	caCertPath := flag.String("ca", defaultCACert, "CA certificate path")
	flag.Parse()

	if *port == defaultPort {
		if envPort := os.Getenv("SENTRY_PORT"); envPort != "" {
			parsedPort, err := strconv.Atoi(envPort)
			if err != nil {
				log.Fatalf("Invalid SENTRY_PORT %q: %v", envPort, err)
			}
			*port = parsedPort
		}
	}
	if *serverCertPath == defaultServerCert {
		if envCert := os.Getenv("SENTRY_SERVER_CERT"); envCert != "" {
			*serverCertPath = envCert
		}
	}
	if *serverKeyPath == defaultServerKey {
		if envKey := os.Getenv("SENTRY_SERVER_KEY"); envKey != "" {
			*serverKeyPath = envKey
		}
	}
	if *caCertPath == defaultCACert {
		if envCA := os.Getenv("SENTRY_CA_CERT"); envCA != "" {
			*caCertPath = envCA
		}
	}

	// Load server certificate and private key
	serverCert, err := tls.LoadX509KeyPair(*serverCertPath, *serverKeyPath)
	if err != nil {
		log.Fatalf("Failed to load server certificates: %v", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile(*caCertPath)
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("sentry.SentryService", grpc_health_v1.HealthCheckResponse_SERVING)
	reflection.Register(s)

	manager := jobmanager.New()
	srv := NewServer(manager)
	pb.RegisterSentryServiceServer(s, srv)

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start signal handler
	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal. Cleaning up...")
		healthServer.Shutdown()
		manager.KillJobsAll()
		os.Exit(0)
	}()

	log.Printf("Server listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
