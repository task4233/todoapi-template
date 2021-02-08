package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/task4233/todoapi-template/internal/db"
)

func TestCreateHandler(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		title    string
		wantCode int
	}{
		"normal": {
			title:    "test",
			wantCode: http.StatusOK,
		},
	}

	db := db.NewMemoryDB()
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			te := CreateRequest{Title: tt.title}
			reqBody, err := json.Marshal(te)
			if err != nil {
				t.Fatalf("failed to json.Marshal: %s\n", err.Error())
			}
			req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/create", bytes.NewBuffer(reqBody))
			got := httptest.NewRecorder()

			handler := &createHandler{
				db: db,
			}
			handler.ServeHTTP(got, req)

			if got.Code != tt.wantCode {
				t.Errorf("expected: %d, actual: %d\n", http.StatusOK, got.Code)
			}

			var resp CreateResponse
			if err := json.NewDecoder(got.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode: %s", err.Error())
			}
			if resp.Title != tt.title {
				t.Errorf("expected: %s, actual: %s", tt.title, resp.Title)
			}
		})
	}
}

type CreateRequest struct {
	Title string `json:"title"`
}

type CreateResponse struct {
	ID    string `json:"ID"`
	Title string `json:"title"`
}

// TODO:
// - testListHandler
// - NewServer
// - Start
// - Stop
