package http_test

import (
	ht "net/http"
	"testing"

	"github.com/task4233/todoapi-template/internal/db"
	"github.com/task4233/todoapi-template/internal/http"
)

type TestNeServer struct {
	port int
	d    db.DB
	want *ht.Server
}

func TestNewServer(t *testing.T) {
	cases := map[string]TestNeServer{
		"normal": {
			port: 8080,
			d:    nil,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			http.NewServer(tt.port, tt.d)
		})
	}
}
