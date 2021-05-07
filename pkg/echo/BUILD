load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["echo.go"],
    importpath = "github.com/neticdk/jytte/pkg/echo",
    visibility = ["//visibility:public"],
    deps = [
        "@io_opentelemetry_go_otel//attribute:go_default_library",
        "@io_opentelemetry_go_otel_trace//:go_default_library",
    ],
)
