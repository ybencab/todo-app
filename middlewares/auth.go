package middlewares

import (
	"net/http"

	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/types"
	"golang.org/x/net/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData := types.UserData{Authenticated: false}

		// Get token from cookie if exists and add user data to context
		cookie, err := r.Cookie("jwt")
		if err == nil {
			claims, err := utils.ValidateJWT(cookie.Value)
			if err == nil {
				userData = types.UserData{
					Authenticated: true,
					ID:      claims["user_id"].(string),
					Email:         claims["email"].(string),
				}
			}
		}

		ctx := context.WithValue(r.Context(), types.UserContextKey, userData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData, ok := utils.GetUserDataFromContext(r.Context())
		if !ok || !userData.Authenticated {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData, ok := utils.GetUserDataFromContext(r.Context())
		if ok && userData.Authenticated {
			http.Redirect(w, r, "/todo", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
