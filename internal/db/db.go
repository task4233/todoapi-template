package db

import (
    "context"
    "github.com/task4233/tododemo/internal/todo"
)

type DB interface {
    PutTODO(ctx context.Context, t *todo.TODO) error
    GetAllTODOs(ctx context.Context) ([]*todo.TODO, error)
}
