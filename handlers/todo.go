package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ybencab/todo-app/store"
	"github.com/ybencab/todo-app/utils"
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
	userData, ok := utils.GetUserDataFromContext(r.Context())
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	todo.Index(&userData).Render(r.Context(), w)
}

func (h *ToDoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {

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
