package main

import (
	"context"
	"fmt"
	ht "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/task4233/todoapi-template/internal/db"
	"github.com/task4233/todoapi-template/internal/http"
)

const addr = 8080

func main() {
	Run()
}

// Run runs server with context
func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	d := db.NewMemoryDB()
	s := http.NewServer(addr, d)
	errCh := make(chan error, 1)

	go func() {
		fmt.Printf("Listening on :%d", addr)
		if err := s.Server.ListenAndServe(); err != nil && err != ht.ErrServerClosed {
			errCh <- fmt.Errorf("failed to server: %w", err)
		}
		errCh <- nil
	}()

	select {
	case <-termCh:
		if err := s.Server.Shutdown(ctx); err != nil {
			errCh <- fmt.Errorf("failed to shutdown: %w", err)
		}
		return 0
	case <-errCh:
		if err := s.Server.Shutdown(ctx); err != nil {
			errCh <- fmt.Errorf("failed to shutdown: %w", err)
		}
		return 1
	}
}
