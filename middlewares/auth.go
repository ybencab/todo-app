package middlewares

import (
	"net/http"

	"github.com/ybencab/todo-app/utils"
	"github.com/ybencab/todo-app/types"
	"golang.org/x/net/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from cookie
		cookie, err := r.Cookie("jwt")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		tokenString := cookie.Value
		
		// Validate token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Get user_id and email from claims
		user_id, ok := claims["user_id"].(string)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		
		// Adding user data to context
		ctx := context.WithValue(
			r.Context(),
			types.UserContextKey,
			types.UserData{
				ID:            user_id,
				Email:         email,
				Authenticated: true,
			},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
