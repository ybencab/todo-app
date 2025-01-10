package handlers

import (
	"net/http"
	"time"

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

	// JWT (For simplicity, we store the token in a cookie non-httponly)
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		components.LoginForm(email, err.Error(), "").Render(r.Context(), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Value: token,
		HttpOnly: false,
		Secure: false,
		SameSite: http.SameSiteStrictMode,
		Path : "/",
		Expires : time.Now().Add(time.Hour * 24),
	})

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}
