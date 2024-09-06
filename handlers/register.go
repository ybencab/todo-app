package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/views/register"
)

func HandleRegisterView(w http.ResponseWriter, r *http.Request) {
	register.Index().Render(r.Context(), w)
}

func HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	user := new(types.CreateUserRequest)
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("error"))
		return
	}

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}