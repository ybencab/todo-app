package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// If user is already logged in, show different header
	userData, ok := utils.GetUserDataFromContext(r.Context())
	if !ok {
		home.Index(nil).Render(r.Context(), w)
		return
	}
	home.Index(&userData).Render(r.Context(), w)
}
