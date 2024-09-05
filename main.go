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
	router.Get("/login", handlers.HandleLogin)
	router.Get("/register", handlers.HandleRegister)

	listenAddr := os.Getenv("LISTEN_ADDR")
	log.Println("HTTP server started in port", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		log.Fatal(err)
	}
}