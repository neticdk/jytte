// Package health provides a simple health endpoint
package health

import "net/http"

type HealthServer struct{}

// NewHandler creates new health server instance
func NewHandler() http.Handler {
	return &HealthServer{}
}

// ServeHTTP request
func (s *HealthServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
