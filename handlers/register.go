package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/store"
	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/validators"
	"github.com/ybencab/todo-app/views/components"
	"github.com/ybencab/todo-app/views/register"
)

type RegisterHandler struct {
	store store.Store
}

func NewRegisterHandler(store store.Store) *RegisterHandler {
	return &RegisterHandler{
		store: store,
	}
}

func (h *RegisterHandler) HandleRegisterView(w http.ResponseWriter, r *http.Request) {
	register.Index().Render(r.Context(), w)
}

func (h *RegisterHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	// TODO: en caso de que la petición venga de HTMX devolveremos el HTML completo del formulario,
	// modificando todo lo que sea necesario del mismo
	// ------
	// Por otro lado, si la petición viene de otro cliente, devoleremos el JSON como en REST API tradicional
	
	// Parse form data
	if err := r.ParseForm(); err != nil {
		components.RegisterForm("", "", "Invalid form data", "").Render(r.Context(), w)
		return
	}
	
	// Getting form values and validating them
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	if err := validators.ValidateRegisterUserRequest(username, email, password); err != nil {
		components.RegisterForm(username, email, err.Error(), "").Render(r.Context(), w)
		return
	}
	
	// Check if user already exists
	if _, err := h.store.GetUserByUsername(username); err == nil {
		components.RegisterForm(username, email, "Username already in use", "").Render(r.Context(), w)
		return
	}
	if _, err := h.store.GetUserByEmail(email); err == nil {
		components.RegisterForm(username, email, "Email already in use", "").Render(r.Context(), w)
		return
	}

	// Create user
	new_user, err := types.NewUser(username, email, password)
	if err != nil {
		components.RegisterForm(username, email, "Error creating user", "").Render(r.Context(), w)
		return
	}

	// Inserting user to database
	if err := h.store.CreateUser(new_user); err != nil {
		components.RegisterForm(username, email, "Error registering user", "").Render(r.Context(), w)
		return
	}

	components.RegisterForm("", "", "", "User registration complete").Render(r.Context(), w)
}
