// Package echo provides a simple echoing http handler
package echo

import (
	"net/http"
)

// EchoServer is a handler implementation echoing back to the requestor
type EchoServer struct {
}

// NewHandler creates new echo server instance
func NewHandler() http.Handler {
	return &EchoServer{}
}

// ServeHTTP request
func (s *EchoServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.Write(w)
}
