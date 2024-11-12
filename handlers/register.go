package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/validators"
	"github.com/ybencab/todo-app/views/register"
)

func HandleRegisterView(w http.ResponseWriter, r *http.Request) {
	register.Index().Render(r.Context(), w)
}

func HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid form data"})
		return
	}
	
	email := r.FormValue("email")
	password := r.FormValue("password")
	if err := validators.ValidateRegisterUserRequest(email, password); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	user := new(types.CreateUserRequest)
	user.Email = email
	user.Password = password

	utils.WriteJSON(w, http.StatusOK, user)
}
