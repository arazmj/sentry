package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pb "github.com/arazmj/sentry-run/api/proto"
	"github.com/arazmj/sentry-run/pkg/jobmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func buildLogger() *slog.Logger {
	var level slog.Level
	switch strings.ToLower(os.Getenv("SENTRY_LOG_LEVEL")) {
	case "debug":
		level = slog.LevelDebug
	case "info", "":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}
	if os.Getenv("SENTRY_LOG_JSON") == "1" {
		return slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}

func main() {
	logger := buildLogger()
	slog.SetDefault(logger)
	jobmanager.SetLogger(logger)

	// Load server certificate and private key
	serverCert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		slog.Error("failed to load server certificates", "error", err)
		os.Exit(1)
	}

	// Load CA certificate
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		slog.Error("failed to load CA certificate", "error", err)
		os.Exit(1)
	}

	// Create certificate pool and append CA certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		slog.Error("failed to append CA certificate to pool")
		os.Exit(1)
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
		slog.Error("failed to listen", "error", err)
		os.Exit(1)
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
		slog.Info("received interrupt signal, cleaning up")
		healthServer.Shutdown()
		manager.KillJobsAll()
		os.Exit(0)
	}()

	slog.Info("server listening", "port", port)
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
		os.Exit(1)
	}
}
