package handlers

import (
	"net/http"

	"github.com/ybencab/todo-app/views/register"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	register.Index().Render(r.Context(), w)
}
