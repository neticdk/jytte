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
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// ListenAndServe instantiates a new server instance
func ListenAndServe() {
	initMetrics()

	if viper.GetBool("tracing") {
		log.Printf("Waiting for tracing connection...")
		shutdown := initTracing()
		defer shutdown()
	}

	http.Handle("/health", health.NewHandler())
	http.Handle("/echo/", otelhttp.NewHandler(echo.NewHandler(), "echo"))
	http.Handle("/entropy/", otelhttp.NewHandler(entropy.NewHandler(), ""))

	log.Printf("Start listening on %s", viper.GetString("listen_address"))
	log.Fatal(http.ListenAndServe(viper.GetString("listen_address"), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
