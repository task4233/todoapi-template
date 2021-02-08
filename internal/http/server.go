package http

import (
    "net/http"
    "fmt"

    "github.com/task4233/tododemo/internal/db"
)

type Server struct {
      server *http.Server
}

func NewServer(port int, d db.DB) *Server{
    mux := http.NewServeMux()
    return &Server {
        server: &http.Server {
            Addr: fmt.Sprintf(":%d", port),
                Handler: mux,
        },
    }
}

func (s *Server) Start() error {
    if s.server.ListenAndServer(); err != nil && err != http.ErrServerClosed {
        return fmt.Errorf("failed to server: %w", err)   
    }

    return nil
}


func (s *Server Stop(ctx context.Context) error {
    if err := s.server.Shutdown(ctx); err != nil {
        return fmt.Errorf("failed to shutdown: %w", err)
    }

    return nil
}
