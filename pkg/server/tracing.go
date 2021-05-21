package server

import (
	"context"
	"time"

	"github.com/neticdk/jytte/pkg/util"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"google.golang.org/grpc"
)

func initTracing(tracingAddr string) func() {
	log.Info().Str("tracingAddr", tracingAddr).Msg("Initialize tracing")

	ctx := context.Background()

	driver := otlpgrpc.NewDriver(
		otlpgrpc.WithInsecure(),
		otlpgrpc.WithEndpoint(tracingAddr),
		otlpgrpc.WithDialOption(grpc.WithBlock(), grpc.WithTimeout(5*time.Second)),
	)
	exp, err := otlp.NewExporter(ctx, driver)
	util.HandleErr(err, "failed to create exporter")

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(ServiceName),
		),
	)
	util.HandleErr(err, "failed to create resource")

	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // sample every trace - scales with load (be careful when using in production)
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	cont := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exp,
		),
		controller.WithExporter(exp),
		controller.WithCollectPeriod(2*time.Second),
	)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)
	global.SetMeterProvider(cont.MeterProvider())
	util.HandleErr(cont.Start(context.Background()), "failed to start controller")

	return func() {
		// Push any last metric events to the exporter.
		util.HandleErr(cont.Stop(context.Background()), "failed to stop controller")

		// Shutdown will flush any remaining spans and shut down the exporter.
		util.HandleErr(tracerProvider.Shutdown(ctx), "failed to shutdown TracerProvider")
	}
}
