workspace(name = "build_stack_subzelle")

load("//:deps.bzl", "base_deps")

base_deps()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("//:go_deps.bzl", "go_deps")

go_deps()
