load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "management.go",
        "plugin.go",
        "server.go",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/logger:go_default_library",
        "//plugin/gcp:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/health:go_default_library",
    ],
)

go_binary(
    name = "idmgrd",
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
    ],
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
