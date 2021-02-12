package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/task4233/todoapi-template/internal/db"
	"github.com/task4233/todoapi-template/internal/http"
)

// Run runs server with context
func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	d := db.NewMemoryDB()
	s := http.NewServer(8080, d)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-time.After(time.Duration(1)):
		if os.Getenv("DEV") == "test" {
			return 0
		}
	case <-termCh:
		return 0
	case <-errCh:
		return 1
	}
	return 0
}
