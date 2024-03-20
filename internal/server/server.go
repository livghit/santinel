package server

import (
	"github.com/gorilla/mux"
	"github.com/livghit/santinel/internal/server/env"
	"github.com/livghit/santinel/internal/server/handlers"
	"github.com/livghit/santinel/internal/server/storage"
	"github.com/livghit/santinel/internal/watchers"
)

// HTTP Server (using Standard Library and Mux)
// the server will be responsible also for frontend
// things like tables , forms and so on via HTMX and ALPINEJS

type Server struct {
	*mux.Router
	Storage *storage.Database
}

func NewServer() *Server {
	// loading the env
	env := env.LoadEnv()

	s := &Server{
		Router:  mux.NewRouter(),
		Storage: storage.New(env),
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	// Here register the routes
	s.HandleFunc("/server/health", handlers.ServerHealthHandler()).Methods("GET")
	s.HandleFunc("/", handlers.DashboardHandler()).Methods("GET")
	s.HandleFunc("/watchers/network", watchers.NetworkHandler).Methods("GET")
	s.Storage.Connection.Ping()
}

func (s *Server) watchers(){
 // here define the watchers we want to use ???
}
