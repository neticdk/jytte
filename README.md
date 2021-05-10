# Jytte

`jytte` is a small http based application written in go with for demo and testing purposes.


## Usage

The application will listen for requests on three endpoints:

`/health`
  will response with a http response `200` when the application is running

`/metrics`
  serves OpenMetrics (Prometheus) endpoint based on OpenTelemetry handler

`/echo/`
  a simple echo server sending any incomming input back to the requestor

`/entropy/`
  a server emulating having multiple dependencies on other services

Besides this the application exposes the Go [pprof](https://golang.org/pkg/net/http/pprof/) allowing runtime profiling and more. This is served on the default path of `/debug/pprof`.


## Development

The project aims at following the best practices as described in [project-layout](https://github.com/golang-standards/project-layout).

Run the application locally simply use `go run` like so:

```bash
go run cmd/jytte/main.go
```

or from Bazel:

```bash
bazel run //cmd/jytte:jytte
```

Building local container image run:

```bash
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/jytte:docker_image
```

Update Bazel dependencies generated from `go.mod`:

```bash
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

Update the Bazel build files:

```bash
bazel run //:gazelle
```
