package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type fakeStoppableServer struct {
	block     bool
	release   chan struct{}
	stopOnce  sync.Once
	stopCalls atomic.Int32
}

func (f *fakeStoppableServer) GracefulStop() {
	if f.block {
		<-f.release
	}
}

func (f *fakeStoppableServer) Stop() {
	f.stopCalls.Add(1)
	f.stopOnce.Do(func() {
		if f.release != nil {
			close(f.release)
		}
	})
}

func TestGracefulShutdownDoesNotForceCompletedStop(t *testing.T) {
	server := &fakeStoppableServer{}
	if graceful := gracefulShutdown(server, 10*time.Millisecond); !graceful {
		t.Fatal("gracefulShutdown() = false, want true")
	}
	time.Sleep(20 * time.Millisecond)

	if calls := server.stopCalls.Load(); calls != 0 {
		t.Fatalf("Stop() calls = %d, want 0", calls)
	}
}

func TestGracefulShutdownForcesBlockedStop(t *testing.T) {
	server := &fakeStoppableServer{
		block:   true,
		release: make(chan struct{}),
	}
	if graceful := gracefulShutdown(server, 10*time.Millisecond); graceful {
		t.Fatal("gracefulShutdown() = true, want false")
	}

	if calls := server.stopCalls.Load(); calls != 1 {
		t.Fatalf("Stop() calls = %d, want 1", calls)
	}
}
