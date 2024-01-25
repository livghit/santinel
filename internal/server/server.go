package server

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// HTTP Server (using Standard Library and Mux)
// the server will be responsible also for frontend
// things like tables , forms and so on via HTMX and ALPINEJS

type Server struct {
	*mux.Router
}

func NewServer(db *gorm.DB) *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	// Here register the routes
}
