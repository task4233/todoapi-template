package server_test

import (
	"testing"

	"github.com/task4233/todoapi-template/internal/server"
)

// TODO:
// - Run
func TestRun(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
	}{
		"normal": {},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			_ = tt
			server.Run()
		})
	}
}
