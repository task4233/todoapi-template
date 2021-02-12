package todo

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TODOs manages an array for todo
type TODOs struct {
	Todos []*TODO `json:"todos"`
}

// TODO mangages information for todo
type TODO struct {
	ID    string     `json:"id"`
	Title string     `json:"title"`
	TS    *time.Time `json:"ts"`
}

// NewTODO creates TODO with uuid.New()
func NewTODO(title string) (*TODO, error) {
	if len(title) == 0 {
		return nil, errors.New("title must not be empty")
	}
	ts := time.Now()
	return &TODO{
		ID:    uuid.New().String(),
		Title: title,
		TS:    &ts,
	}, nil
}

// IsValid checks if TODO is a valid
// if valid, returns nil
func (t TODO) IsValid() error {
	if len(t.ID) == 0 {
		return errors.New("ID must not be empty")
	}

	if len(t.Title) == 0 {
		return errors.New("title must not be empty")
	}

	if t.TS == nil {
		ts := time.Now()
		t.TS = &ts
	}
	return nil
}
