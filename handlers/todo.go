package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ybencab/todo-app/store"
	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/validators"
	"github.com/ybencab/todo-app/views/components"
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

	todos, err := h.store.GetTodos(userData.ID)
	if err != nil {
		todo.Index(&userData, nil).Render(r.Context(), w)
		return
	}

	todo.Index(&userData, todos).Render(r.Context(), w)
}

func (h *ToDoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		components.ToDoForm("", "", "Invalid form data", "").Render(r.Context(), w)
		return
	}

	// Getting form values and validate them
	title := r.FormValue("title")
	description := r.FormValue("description")
	if err := validators.ValidateCreateTodoRequest(title, description); err != nil {
		components.ToDoForm(title, description, err.Error(), "").Render(r.Context(), w)
		return
	}

	// Get user data
	userData, ok := utils.GetUserDataFromContext(r.Context())
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if todo already exists
	if _, err := h.store.GetTodo(userData.ID, title); err == nil {
		components.ToDoForm(title, description, "Todo already exists", "").Render(r.Context(), w)
		return
	}

	// Convert user_id to uuid
	user_id, err := uuid.Parse(userData.ID)
	if err != nil {
		components.ToDoForm(title, description, "Error creating todo", "").Render(r.Context(), w)
		return
	}

	// Create todo
	new_todo := types.NewTodo(user_id, title, description)
	if err := h.store.CreateTodo(new_todo); err != nil {
		components.ToDoForm(title, description, "Error creating todo", "").Render(r.Context(), w)
		return
	}
	
	w.Header().Set("Hx-Redirect", "/todo")
}

func (h *ToDoHandler) HandletGetTodos(w http.ResponseWriter, r *http.Request) {
	userData, ok := utils.GetUserDataFromContext(r.Context())
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	todos, err := h.store.GetTodos(userData.ID)
	if err != nil {
		http.Error(w, "Error getting todos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
