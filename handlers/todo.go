package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/views/todo"
)

func HandleTodo(w http.ResponseWriter, r *http.Request) {
	todo.Index().Render(r.Context(), w)
}

func HandletGetTodo(w http.ResponseWriter, r *http.Request) {

}
