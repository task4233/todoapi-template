package todo_test

import (
	"testing"

	"github.com/task4233/todoapi-template/internal/todo"
)

func TestNewTODO(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		title   string
		want    *todo.TODO
		wantErr string
	}{
		"nomal": {
			title: "test",
			want:  &todo.TODO{Title: "test"},
		},
		"empty title": {
			wantErr: "title must not be empty",
		},
	}

	for name, tc := range cases {
		test := tc
		t.Run(name, func(t *testing.T) {
			td, err := todo.NewTODO(tc.title)
			if err != nil {
				if err.Error() != test.wantErr {
					t.Errorf("expected: %s, actual: %s\n", test.wantErr, err.Error())
				}
			} else if err := td.IsValid(); err != nil {
				if err.Error() != test.wantErr {
					t.Errorf("expected: %s, actual: %s\n", test.wantErr, err)
				}
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		arg     *todo.TODO
		wantErr string
	}{
		"nomal": {
			arg: &todo.TODO{ID: "test", Title: "test"},
		},
		"empty title": {
			arg:     &todo.TODO{ID: "test"},
			wantErr: "title must not be empty",
		},
		"empty ID": {
			arg:     &todo.TODO{Title: "test"},
			wantErr: "ID must not be empty",
		},
	}

	for name, tc := range cases {
		test := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if err := test.arg.IsValid(); err != nil {
				if err.Error() != test.wantErr {
					t.Errorf("expected: %s, actual: %s\n", test.wantErr, err)
				}
			}
		})
	}
}
