load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["logger.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/grpc/grpclog/glogger:go_default_library"],
)
