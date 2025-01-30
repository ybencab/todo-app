package store

import (
	"github.com/ybencab/todo-app/types"
)

type Store interface {
	CreateTodo(*types.ToDo) error
	GetTodo(string, string) (*types.ToDo, error)
	GetTodos(string) ([]*types.ToDo, error)
	CreateUser(*types.User) error
	GetUserByUsername(string) (*types.User, error)
	GetUserByEmail(string) (*types.User, error)
	GetAllUsers() ([]*types.User, error)
}
