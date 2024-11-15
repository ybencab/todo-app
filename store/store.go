package store

import (
	"github.com/ybencab/todo-app/types"
)

type Store interface {
	CreateTodo(*types.ToDo) error
	GetTodo(int) (*types.ToDo, error)
	GetTodos() ([]*types.ToDo, error)
	CreateUser(*types.User) error
	GetUserByUsername(string) (*types.User, error)
	GetUserByEmail(string) (*types.User, error)
	GetAllUsers() ([]*types.User, error)
}
