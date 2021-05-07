// Package echo provides a simple echoing http handler
package echo

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	apiTrace "go.opentelemetry.io/otel/trace"
)

// Server is a handler implementation echoing back to the requestor
type Server struct {
}

// NewHandler creates new echo server instance
func NewHandler() http.Handler {
	return &Server{}
}

// ServeHTTP request
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	span := apiTrace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("foo", "bar"))
	req.Write(w)
}
