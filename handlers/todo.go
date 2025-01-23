package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ybencab/todo-app/store"
	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/views/todo"
)

type ToDoHandler struct {
	store store.Store
}

func NewTodoHandler(store store.Store) *ToDoHandler {
	return &ToDoHandler{
		store: store,
	}
}

func (h *ToDoHandler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	todo.Index(r).Render(r.Context(), w)
}

func (h *ToDoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("error"))
		return
	}

	todoTitle := r.FormValue("title")
	todoDescription := r.FormValue("description")
	todo := types.NewTodo(todoTitle, todoDescription)

	if err := h.store.CreateTodo(todo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(todo)
}

func (h *ToDoHandler) HandletGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetTodos()
	if err != nil {
		w.Write([]byte("error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(todos)
}
