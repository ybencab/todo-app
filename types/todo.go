package types

import (
	"time"
)

type ToDo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(title, description string) *ToDo {
	return &ToDo{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
