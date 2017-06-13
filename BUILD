package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@io_kubernetes_build//defs:build.bzl", "gcs_upload")

go_prefix("containers.google.com/idmgr")

gcs_upload(
    name = "push-build",
    data = [
        "//cmd/idmgr_cli",
        "//cmd/idmgrd",
    ],
)
