package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/store"
	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/views/components"
	"github.com/ybencab/todo-app/views/login"
)

type LoginHandler struct {
	store store.Store
}

func NewLoginHandler(store store.Store) *LoginHandler {
	return &LoginHandler{
		store: store,
	}
}

func (h *LoginHandler) HandleLoginView(w http.ResponseWriter, r *http.Request) {
	login.Index().Render(r.Context(), w)
}

func (h *LoginHandler) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		components.LoginForm("", "Invalid form data", "").Render(r.Context(), w)
		return
	}

	// Getting form values
	email := r.FormValue("email")
	password := r.FormValue("password")


	// Check if user exists
	user, err := h.store.GetUserByEmail(email)
	if err != nil {
		components.LoginForm(email, "User not found", "").Render(r.Context(), w)
		return
	}

	// Compare password
	if err := utils.CompareHashAndPassword(user.Password, password); err != nil {
		components.LoginForm(email, "Invalid password", "").Render(r.Context(), w)
		return
	}

	// Success message
	components.LoginForm(email, "Login successful", "").Render(r.Context(), w)
}
