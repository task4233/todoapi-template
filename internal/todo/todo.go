package todo

import (
	"errors"

	"github.com/google/uuid"
)

// TODOs manages an array for todo
type TODOs struct {
	Todos []*TODO `json:"todos"`
}

// TODO mangages information for todo
type TODO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// NewTODO creates TODO with uuid.New()
func NewTODO(title string) (*TODO, error) {
	if len(title) == 0 {
		return nil, errors.New("title must not be empty")
	}

	return &TODO{
		ID:    uuid.New().String(),
		Title: title,
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
	return nil
}
