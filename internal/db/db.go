package db

import (
    "context"
    "github.com/task4233/tododemo/internal/todo"
)

type DB interface {
    putTODO(ctx context.Context, t *todo.TODO) error
}
