load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["entropy.go"],
    importpath = "github.com/neticdk/jytte/pkg/entropy",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_rs_zerolog//log:go_default_library",
        "@io_opentelemetry_go_otel//:go_default_library",
        "@io_opentelemetry_go_otel//attribute:go_default_library",
        "@io_opentelemetry_go_otel//codes:go_default_library",
        "@io_opentelemetry_go_otel_metric//:go_default_library",
        "@io_opentelemetry_go_otel_metric//global:go_default_library",
        "@io_opentelemetry_go_otel_trace//:go_default_library",
    ],
)
