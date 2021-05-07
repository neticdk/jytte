package server

import (
	"log"
	"net/http"

	"go.opentelemetry.io/otel/exporters/metric/prometheus"
)

func initMetrics() {
	exporter, err := prometheus.InstallNewPipeline(prometheus.Config{})
	if err != nil {
		log.Panicf("failed to initialize prometheus exporter %v", err)
	}
	http.HandleFunc("/metrics", exporter.ServeHTTP)
}
