package server

import (
	"github.com/gorilla/mux"
	"github.com/livghit/santinel/internal/server/database"
	"github.com/livghit/santinel/internal/server/env"
	"github.com/livghit/santinel/internal/server/handlers"
)

// HTTP Server (using Standard Library and Mux)
// the server will be responsible also for frontend
// things like tables , forms and so on via HTMX and ALPINEJS

type Server struct {
	*mux.Router
	*database.Database
}

func NewServer() *Server {
	// loading the env
	env := env.LoadEnv()

	s := &Server{
		Router:   mux.NewRouter(),
		Database: database.New(env),
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	// Here register the routes
	s.HandleFunc("/", handlers.TestHandler()).Methods("GET")
}
