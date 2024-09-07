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
	Store  store.Store
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL not defined")
	}

	store := store.NewPostgresStore(connStr)

	server := NewServer(store)
	server.MountHandlers()

	listenAddr := os.Getenv("PORT")
	if listenAddr == "" {
		log.Fatal("PORT not defined")
	}

	log.Println("HTTP server started in port", listenAddr)
	if err := http.ListenAndServe(listenAddr, server.Router); err != nil {
		log.Fatal(err)
	}
}

func NewServer(store store.Store) *Server {
	return &Server{
		Router: chi.NewRouter(),
		Store:  store,
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
