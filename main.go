package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/ybencab/todo-app/handlers"
	"github.com/ybencab/todo-app/store"
)

type Server struct {
	Router *chi.Mux
	Store  *store.Store
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	server := NewServer()
	server.MountHandlers()

	listenAddr := os.Getenv("LISTEN_ADDR")
	log.Println("HTTP server started in port", listenAddr)
	if err := http.ListenAndServe(listenAddr, server.Router); err != nil {
		log.Fatal(err)
	}
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) MountHandlers() {
	s.Router.Handle("/*", public())

	s.Router.Get("/", handlers.HandleHome)
	s.Router.Route("/login", func(r chi.Router) {
		r.Get("/", handlers.HandleLoginView)
		r.Post("/", handlers.HandleLoginUser)
	})
	s.Router.Route("/register", func(r chi.Router) {
		r.Get("/", handlers.HandleRegisterView)
		r.Post("/", handlers.HandleRegisterUser)
	})
}
