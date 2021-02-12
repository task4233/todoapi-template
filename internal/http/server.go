package http

import (
	"fmt"
	"net/http"

	"github.com/task4233/todoapi-template/internal/db"
)

// Server is a struct for *http.Server
type Server struct {
	Server *http.Server
}

// NewServer provides a pointer to Server instance
func NewServer(port int, d db.DB) *Server {
	mux := http.NewServeMux()

	mux.Handle("/create", &createHandler{db: d})
	mux.Handle("/list", &listHandler{db: d})

	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}
