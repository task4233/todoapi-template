package db

import (
	"context"

	"github.com/task4233/todoapi-template/internal/todo"
)

// DB is an interface for Database
type DB interface {
	PutTODO(ctx context.Context, t *todo.TODO) error
	GetAllTODOs(ctx context.Context) (*todo.TODOs, error)
}
