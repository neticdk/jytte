// Package health provides a simple health endpoint
package health

import "net/http"

// Server is a handler implementation
type Server struct{}

// NewHandler creates new health server instance
func NewHandler() http.Handler {
	return &Server{}
}

// ServeHTTP request
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
