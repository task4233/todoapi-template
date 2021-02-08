package db

import (
    "context"
    "sync"

    "github.com/task4233/tododemo/internal/todo"
)

var _ DB = (*memoryDB)(nil)

type memoryDB struct {
    db map[string]*todo.TODO
    lock sync.RWMutex
    
}

func (m *memoryDB) PutTODO(ctx context.Context, t *todo.TODO) error {
    m.lock.Lock()
    m.db[t.ID] = t
    m.lock.Unlock()
    return nil
}
