package types

import (
	"time"

	"github.com/google/uuid"
)

type ToDo struct {
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(user_id uuid.UUID, title, description string) *ToDo {
	return &ToDo{
		UserID: user_id,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
