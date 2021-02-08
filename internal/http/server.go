package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/task4233/todoapi-template/internal/db"
)

// Server is a struct for *http.Server
type Server struct {
	server *http.Server
}

// NewServer provides a pointer to Server instance
func NewServer(port int, d db.DB) *Server {
	mux := http.NewServeMux()

	mux.Handle("/create", &createHandler{db: d})
	mux.Handle("/list", &listHandler{db: d})

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

// Start starts server
func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to server: %w", err)
	}

	return nil
}

// Stop stops server
func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}
