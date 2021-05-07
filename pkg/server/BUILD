load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "metrics.go",
        "server.go",
        "tracing.go",
    ],
    importpath = "github.com/neticdk/jytte/pkg/server",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/echo:go_default_library",
        "//pkg/entropy:go_default_library",
        "//pkg/health:go_default_library",
        "//pkg/util:go_default_library",
        "@com_github_gorilla_handlers//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
        "@io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp//:go_default_library",
        "@io_opentelemetry_go_otel//:go_default_library",
        "@io_opentelemetry_go_otel//propagation:go_default_library",
        "@io_opentelemetry_go_otel//semconv:go_default_library",
        "@io_opentelemetry_go_otel_exporters_metric_prometheus//:go_default_library",
        "@io_opentelemetry_go_otel_exporters_otlp//:go_default_library",
        "@io_opentelemetry_go_otel_exporters_otlp//otlpgrpc:go_default_library",
        "@io_opentelemetry_go_otel_metric//global:go_default_library",
        "@io_opentelemetry_go_otel_sdk//resource:go_default_library",
        "@io_opentelemetry_go_otel_sdk//trace:go_default_library",
        "@io_opentelemetry_go_otel_sdk_metric//controller/basic:go_default_library",
        "@io_opentelemetry_go_otel_sdk_metric//processor/basic:go_default_library",
        "@io_opentelemetry_go_otel_sdk_metric//selector/simple:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
