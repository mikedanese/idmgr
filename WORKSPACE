http_archive(
    name = "io_bazel_rules_go",
    sha256 = "",
    strip_prefix = "rules_go-936af5753ebcd7a1f05127678435389cc2e3db5d",
    urls = ["https://github.com/bazelbuild/rules_go/archive/936af5753ebcd7a1f05127678435389cc2e3db5d.tar.gz"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")

go_proto_repositories()

http_archive(
    name = "io_kubernetes_build",
    sha256 = "a9fb7027f060b868cdbd235a0de0971b557b9d26f9c89e422feb80f48d8c0e90",
    strip_prefix = "repo-infra-9dedd5f4093884c133ad5ea73695b28338b954ab",
    urls = ["https://github.com/kubernetes/repo-infra/archive/9dedd5f4093884c133ad5ea73695b28338b954ab.tar.gz"],
)
