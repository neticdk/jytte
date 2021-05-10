// Package server provides the main server implementation
package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/neticdk/jytte/pkg/echo"
	"github.com/neticdk/jytte/pkg/entropy"
	"github.com/neticdk/jytte/pkg/health"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// ServiceName is used to annotate traces and metrics
const ServiceName string = "jytte"

// ListenAndServe instantiates a new server instance
func ListenAndServe(listenAddr string, tracingEnabled bool, tracingAddr string) {
	if tracingEnabled {
		log.Printf("Waiting for tracing connection...")
		shutdown := initTracing(tracingAddr)
		defer shutdown()
	}

	initMetrics()

	http.Handle("/health", health.NewHandler())
	http.Handle("/echo/", otelhttp.NewHandler(echo.NewHandler(), "echo"))
	http.Handle("/entropy/", otelhttp.NewHandler(entropy.NewHandler(), ""))

	log.Printf("Start listening on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
