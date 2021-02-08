package db

import (
    "context"
    "testing"

    "github.com/google/go-cmp/cmp"
   
    "github.com/task4233/tododemo/internal/todo"
)

func TestMemoryDBPutTODO(t *testing.T) {
    t.Parallel()

    todo1 := &todo.TODO{
        ID: "test",
        Title: "testTitle",
    }

    cases := map[string]struct{
        todo *todo.TODO
        want map[string]*todo.TODO
    }{
        "put": {
            todo: todo1,
            want: map[string]*todo.TODO{todo1.ID: todo1}, 
        },
    }

    ctx := context.Background()
    for name, tc := range cases {
        test := tc
        t.Run(name, func(t *testing.T) {
            t.Parallel()

            d := &memoryDB{db: map[string]*todo.TODO{}}
            if err := d.PutTODO(ctx, test.todo); err != nil {
                t.Fatalf("failed to put a todo: %s", err.Error())
            }

            if diff := cmp.Diff(test.want, d.db); diff != "" {
                t.Errorf("\n(-expected, +actual)\n%s", diff)
            }
        })
    }
}
