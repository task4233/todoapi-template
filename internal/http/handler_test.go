package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/task4233/todoapi-template/internal/db"
	"github.com/task4233/todoapi-template/internal/todo"
)

type failDB struct{}

func NewFailDB() db.DB {
	return &failDB{}
}

func (f failDB) PutTODO(ctx context.Context, t *todo.TODO) error {
	return errors.New("fails intentionally")
}

func (f failDB) GetAllTODOs(ctx context.Context) (*todo.TODOs, error) {
	return nil, errors.New("fails intentionally")
}

type testStructPutTODO struct {
	db        db.DB
	reqBody   []byte
	wantTitle string
	wantCode  int
}

func TestCreateHandler(t *testing.T) {
	var title string = "normal"
	te := CreateRequest{Title: title}
	reqBody, err := json.Marshal(te)
	if err != nil {
		t.Fatalf("failed to json.Marshal: %s\n", err.Error())
	}

	db := db.NewMemoryDB()
	failDB := NewFailDB()

	cases := map[string]testStructPutTODO{
		"normal": {
			reqBody:   reqBody,
			db:        db,
			wantTitle: title,
			wantCode:  http.StatusOK,
		},
		"invalid Body": {
			reqBody:   nil,
			db:        db,
			wantTitle: "",
			wantCode:  http.StatusBadRequest,
		},
		"failed PutTODO": {
			reqBody:   reqBody,
			db:        failDB,
			wantTitle: "",
			wantCode:  http.StatusInternalServerError,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/create", bytes.NewBuffer(tt.reqBody))
			got := httptest.NewRecorder()

			handler := &createHandler{
				db: tt.db,
			}
			handler.ServeHTTP(got, req)

			if got.Code != tt.wantCode {
				t.Errorf("expected: %d, actual: %d\n", http.StatusOK, got.Code)
			}

			// StatusOKの時のみ中身を検証する
			if got.Code != http.StatusOK {
				return
			}

			var resp CreateResponse
			if err := json.NewDecoder(got.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode: %s", err.Error())
			}
			if resp.Title != tt.wantTitle {
				t.Errorf("expected: %s, actual: %s", tt.wantTitle, resp.Title)
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

type TestStructGetAllTODOs struct {
	db       db.DB
	want     []todo.TODO
	wantCode int
}

func TestGetAllHandler(t *testing.T) {
	db := db.NewMemoryDB()
	ctx := context.Background()
	td, err := todo.NewTODO("normal")
	if err != nil {
		t.Fatalf("failed to NewTODO: %s", err.Error())
	}

	if err := db.PutTODO(ctx, td); err != nil {
		t.Fatalf("failed to PutTODO: %s", err.Error())
	}
	failDB := NewFailDB()

	cases := map[string]TestStructGetAllTODOs{
		"normal": {
			db:       db,
			want:     []todo.TODO{},
			wantCode: http.StatusOK,
		},
		"failed PutGetAllTODOs": {
			db:       failDB,
			want:     nil,
			wantCode: http.StatusInternalServerError,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/list", nil)
			got := httptest.NewRecorder()

			handler := &listHandler{
				db: tt.db,
			}
			handler.ServeHTTP(got, req)

			if got.Code != tt.wantCode {
				t.Errorf("expected: %d, actual: %d\n", http.StatusOK, got.Code)
			}

			// StatusOKの時のみ中身を検証する
			if got.Code != http.StatusOK {
				return
			}

			log.Println("body:", got.Body)
			var resp GetAllResponse
			if err := json.NewDecoder(got.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode: %s", err.Error())
			}
		})
	}
}

type GetAllResponse struct {
	Todos []CreateResponse `json:"todos"`
}

// TODO:
// - testListHandler
// - NewServer
// - Start
// - Stop
