load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/neticdk/jytte/cmd/jytte",
    visibility = ["//visibility:private"],
    deps = ["//pkg/cmd:go_default_library"],
)

go_binary(
    name = "jytte",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

container_image(
    name = "docker_image",
    base = "@alpine_linux_amd64//image",
    entrypoint = ["/jytte"],
    files = [":jytte"],
    user = "65534",
)

container_push(
    name = "publish",
    format = "Docker",
    image = ":docker_image",
    registry = "ghcr.io",
    repository = "neticdk/jytte",
    tag = "$(version)",
)
