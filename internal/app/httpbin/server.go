package httpbin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents the server
type Server struct {
	router *mux.Router
}

// NewServer builds and returns a new server
func NewServer(router *mux.Router) (*Server, error) {
	server := &Server{
		router: router,
	}
	server.initRoutes() // register handlers
	return server, nil
}

// ListenAndServe starts the http listener
func (s *Server) ListenAndServe() {
	log.Fatal(http.ListenAndServe(":8080", s.router))
}

