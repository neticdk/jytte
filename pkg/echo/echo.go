// Package echo provides a simple echoing http handler
package echo

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel/trace"
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
	span := trace.SpanFromContext(req.Context())
	log.Info().Str("TraceID", span.SpanContext().TraceID().String()).Msg("Echo service")

	req.Write(w)
}
