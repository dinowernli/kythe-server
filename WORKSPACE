workspace(name = "me_dinowernli_faucet")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "afec53d875013de6cebe0e51943345c587b41263fdff36df5ff651fbf03c1c08",
    strip_prefix = "rules_go-0.4.4",
    url = "https://github.com/bazelbuild/rules_go/archive/0.4.4.tar.gz",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")
load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")

go_repositories()

go_proto_repositories()

http_archive(
    name = "io_bazel",
    sha256 = "8e4646898fa9298422e69767752680d34cbf21bcae01c401b11aa74fcdb0ef66",
    strip_prefix = "bazel-0.4.1",
    url = "https://github.com/bazelbuild/bazel/archive/0.4.1.tar.gz",
)

new_go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/Sirupsen/logrus",
    tag = "v0.11.0",
)

