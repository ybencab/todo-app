package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ybencab/todo-app/types"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

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
