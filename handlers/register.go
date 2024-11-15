package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/validators"
	"github.com/ybencab/todo-app/views/register"
	"github.com/ybencab/todo-app/views/components"
)

func HandleRegisterView(w http.ResponseWriter, r *http.Request) {
	register.Index().Render(r.Context(), w)
}

func HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	// TODO: en caso de que la petición venga de HTMX devolveremos el HTML completo del formulario,
	// modificando todo lo que sea necesario del mismo
	// ------
	// Por otro lado, si la petición viene de otro cliente, devoleremos el JSON como en REST API tradicional
	if err := r.ParseForm(); err != nil {
		components.RegisterForm("", "Invalid form data", "").Render(r.Context(), w)
		return
	}
	
	email := r.FormValue("email")
	password := r.FormValue("password")
	if err := validators.ValidateRegisterUserRequest(email, password); err != nil {
		components.RegisterForm(email, err.Error(), "").Render(r.Context(), w)
		return
	}
	
	user := new(types.CreateUserRequest)
	user.Email = email
	user.Password = password

	components.RegisterForm("", "", user.Email + " " + user.Password).Render(r.Context(), w)
}
