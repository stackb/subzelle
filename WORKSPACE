workspace(name = "build_stack_subzelle")

load("//:deps.bzl", "base_deps")

base_deps()

########### Protobuf dependencies #######

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

########### gRPC dependencies #######

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

load("//:bazel_version_repository.bzl", "bazel_version_repository")

bazel_version_repository(
    name = "upb_bazel_version",
)

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

########### Go/Gazelle dependencies #######

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

load("//:go_deps.bzl", "go_deps")

go_deps()
