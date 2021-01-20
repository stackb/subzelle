load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_deps():
    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:zvIju4sqAGvwKspUQOhwnpcqSbzi7/H6QomNNjTL4sk=",
        version = "v1.27.1",
    )

    go_repository(
        name = "org_golang_x_net",
        commit = "a04bdaca5b32abe1c069418fb7088ae607de5bd0",  # master as of 2017-10-10
        importpath = "golang.org/x/net",
    )

    go_repository(
        name = "org_golang_x_text",
        commit = "ab5ac5f9a8deb4855a60fab02bc61a4ec770bd49",  # v0.1.0, latest as of 2017-10-10
        importpath = "golang.org/x/text",
    )

    go_repository(
        name = "org_golang_google_genproto",
        commit = "f676e0f3ac6395ff1a529ae59a6670878a8371a6",  # master on 2017-10-10
        importpath = "google.golang.org/genproto",
    )
