package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/views/login"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	login.Index().Render(r.Context(), w)
}
