load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    importpath = "github.com/neticdk/jytte/pkg/server",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/echo:go_default_library",
        "//pkg/health:go_default_library",
        "@com_github_gorilla_handlers//:go_default_library",
        "@io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp//:go_default_library",
        "@io_opentelemetry_go_otel_exporters_metric_prometheus//:go_default_library",
    ],
)
