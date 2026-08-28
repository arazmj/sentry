package main

import (
	"log/slog"
	"time"
)

const gracefulShutdownTimeout = 30 * time.Second

type stoppableServer interface {
	GracefulStop()
	Stop()
}

func gracefulShutdown(server stoppableServer, timeout time.Duration) bool {
	forced := make(chan struct{})
	forceStop := time.AfterFunc(timeout, func() {
		close(forced)
		slog.Warn("graceful shutdown timed out; forcing gRPC server stop", "timeout", timeout)
		server.Stop()
	})

	server.GracefulStop()
	if forceStop.Stop() {
		return true
	}
	<-forced
	return false
}
