package utils

import (
	"database/sql"

	"github.com/ybencab/todo-app/types"
)

func ScanRowIntoToDo(row *sql.Rows) (*types.ToDo, error) {
	todo := new(types.ToDo)
	err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.CreatedAt,
	)
	return todo, err
}
