package handlers

import (
	"net/http"
	"time"
)

type LogoutHandler struct{}

func NewLogoutHandler() *LogoutHandler {
	return &LogoutHandler{}
}

func (h *LogoutHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Delete cookie
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
