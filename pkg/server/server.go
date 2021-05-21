// Package server provides the main server implementation
package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/neticdk/jytte/pkg/echo"
	"github.com/neticdk/jytte/pkg/entropy"
	"github.com/neticdk/jytte/pkg/health"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// ServiceName is used to annotate traces and metrics
const ServiceName string = "jytte"

// ListenAndServe instantiates a new server instance
func ListenAndServe(listenAddr string, tracingEnabled bool, tracingAddr string) {
	if tracingEnabled {
		log.Info().Msg("Waiting for tracing connection to start up...")
		shutdown := initTracing(tracingAddr)
		defer shutdown()
	}

	initMetrics()

	http.Handle("/health", health.NewHandler())
	http.Handle("/echo/", otelhttp.NewHandler(echo.NewHandler(), "echo"))
	http.Handle("/entropy/", otelhttp.NewHandler(entropy.NewHandler(), ""))

	log.Info().Str("listenAddr", listenAddr).Msg("Start listening")
	log.Fatal().Err(http.ListenAndServe(listenAddr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))).Msg("Server failed")
}
