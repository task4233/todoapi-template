package http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/task4233/todoapi-template/internal/db"
	"github.com/task4233/todoapi-template/internal/todo"
)

var _ http.Handler = (*createHandler)(nil)
var _ http.Handler = (*listHandler)(nil)

type createHandler struct {
	db db.DB
}

// ServeHTTP is for createHandler, POST /create
func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t todo.TODO
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ID = uuid.New().String()
	if err := h.db.PutTODO(r.Context(), &t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type listHandler struct {
	db db.DB
}

func (h *listHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := h.db.GetAllTODOs(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if json.NewEncoder(w).Encode(&todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
