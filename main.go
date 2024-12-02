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
	log.Println("Starting server...")
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	log.Println("Connecting to database...")
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL not defined")
	}

	store := store.NewPostgresStore(connStr)
	if err := store.Init(); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

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

	todoHandler := handlers.NewTodoHandler(s.Store)
	loginHandler := handlers.NewLoginHandler(s.Store)
	registerHandler := handlers.NewRegisterHandler(s.Store)

	s.Router.Get("/", handlers.HandleHome)
	s.Router.Route("/login", func(r chi.Router) {
		r.Get("/", loginHandler.HandleLoginView)
		r.Post("/", loginHandler.HandleLoginUser)
	})
	s.Router.Route("/register", func(r chi.Router) {
		r.Get("/", registerHandler.HandleRegisterView)
		r.Post("/", registerHandler.HandleRegisterUser)
	})
	s.Router.Route("/todo", func(r chi.Router) {
		r.Get("/", todoHandler.HandleTodo)
		r.Post("/", todoHandler.HandleCreateTodo)
		r.Get("/all", todoHandler.HandletGetTodos)
	})
}
