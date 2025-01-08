package middlewares

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated
		// If not, redirect to login page
		// If yes, call next handler
		
		next.ServeHTTP(w, r)
	})
}
