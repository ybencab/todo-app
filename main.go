package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/ybencab/todo-app/handlers"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	
	router := chi.NewRouter()

	router.Handle("/*", public())

	router.Get("/", handlers.HandleHome)
	
	router.Route("/login", func(r chi.Router) {
		r.Get("/", handlers.HandleLoginView)
		r.Post("/", handlers.HandleLoginUser)
	})
	
	router.Route("/register", func(r chi.Router) {
		r.Get("/", handlers.HandleRegisterView)
		r.Post("/", handlers.HandleRegisterUser)
	})

	listenAddr := os.Getenv("LISTEN_ADDR")
	log.Println("HTTP server started in port", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		log.Fatal(err)
	}
}