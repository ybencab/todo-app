package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	home.Index(r).Render(r.Context(), w)
}
