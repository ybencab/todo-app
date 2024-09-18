package store

import "github.com/ybencab/todo-app/types"

type Store interface {
	CreateTodo(*types.ToDo) error
	GetTodo(int) (*types.ToDo, error)
	GetTodos() ([]*types.ToDo, error)
}
