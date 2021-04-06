// Package server provides the main server implementation
package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/neticdk/jytte/pkg/echo"
	"github.com/neticdk/jytte/pkg/health"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/exporters/metric/prometheus"
)

// ListenAndServe instantiates a new server instance
func ListenAndServe(addr string) {
	initMetrics()
	initTracing()

	http.Handle("/health", health.NewHandler())
	http.Handle("/echo/", otelhttp.NewHandler(echo.NewHandler(), "echo"))

	log.Printf("Start listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

func initMetrics() {
	exporter, err := prometheus.InstallNewPipeline(prometheus.Config{})
	if err != nil {
		log.Panicf("failed to initialize prometheus exporter %v", err)
	}
	http.HandleFunc("/metrics", exporter.ServeHTTP)
}

func initTracing() {

}
