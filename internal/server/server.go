package server

import (
	"github.com/gorilla/mux"
	"github.com/livghit/santinel/internal/server/handlers"
	"gorm.io/gorm"
)

// HTTP Server (using Standard Library and Mux)
// the server will be responsible also for frontend
// things like tables , forms and so on via HTMX and ALPINEJS

type Server struct {
	*mux.Router
	*gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		DB:     db,
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	// Here register the routes
	s.HandleFunc("/", handlers.TestHandler()).Methods("GET")
}
