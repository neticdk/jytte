// Package entropy provides a http handler that introduces random delays
package entropy

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/trace"
)

// Server is a handler implementation
type Server struct {
}

// NewHandler creates new entropy server instance
func NewHandler() http.Handler {
	return &Server{}
}

// ServeHTTP request
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tracer := otel.Tracer("entropy-tracer")

	commonLabels := []attribute.KeyValue{
		attribute.String("endpoint", "entropy"),
	}

	rand.Seed(time.Now().UnixNano())

	ctx, span := tracer.Start(
		req.Context(),
		"Call-Backend-Mock-Services",
		trace.WithAttributes(commonLabels...))
	defer span.End()

	meter := global.GetMeterProvider().Meter("entropy")
	errorUpDownCounter := metric.Must(meter).NewInt64UpDownCounter("entropy.backend.errors")
	latencyValueRecorder := metric.Must(meter).NewInt64ValueRecorder("entropy.backend.latency")

	// holds the output sent back to the HTTP client
	var responseText string
	// set to true if any of the mocked backend services fail - used to make the whole request fail
	var backendServiceFailed bool

	for i := 1; i < 11; i++ {
		target := fmt.Sprintf("/service-%d", i)
		_, iSpan := tracer.Start(
			ctx,
			target,
			trace.WithSpanKind(trace.SpanKindClient))

		backendLabels := []attribute.KeyValue{
			attribute.String("http.flavor", "1.1"),
			attribute.String("http.host", "localhost:8080"),
			attribute.String("http.method", "GET"),
			attribute.String("http.scheme", "http"),
			attribute.String("http.target", target),
		}

		metricLabels := []attribute.KeyValue{
			attribute.String("target", target),
		}

		// Generate a status code
		statusCode := rand.Intn(10) + 1

		statusText := "success"
		httpStatusCode := 200
		spanStatusCode := codes.Ok

		// Make 20% of requests fail
		if statusCode < 3 {
			httpStatusCode = 500
			statusText = "failed"
			spanStatusCode = codes.Error
			backendServiceFailed = true
			errorUpDownCounter.Add(ctx, 1, metricLabels...)
		}

		iSpan.SetStatus(spanStatusCode, statusText)
		backendLabels = append(backendLabels, attribute.Int("http.status_code", httpStatusCode))
		responseText += fmt.Sprintf(
			"target:%s i:%d total:10 status:%d status_text:%s\n",
			target, i, statusCode, statusText)

		iSpan.SetAttributes(backendLabels...)

		latency := rand.Intn(500-10) + 10
		<-time.After(time.Duration(latency) * time.Millisecond)
		latencyValueRecorder.Record(ctx, int64(latency), metricLabels...)
		iSpan.End()
	}

	if backendServiceFailed {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(responseText))
	req.Write(w)
}
