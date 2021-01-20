load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def base_deps():
    http_archive(
        name = "com_google_protobuf",
        sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
        strip_prefix = "protobuf-3.14.0",
        urls = [
            "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        ],
    )

    http_archive(
        name = "com_github_grpc_grpc",
        sha256 = "c260a1dcdd26a78a9596494a3f41f9594ab5ec3a4d65cba4658bdee2b55ac844",
        strip_prefix = "grpc-1.34.1",
        urls = [
            "https://github.com/grpc/grpc/archive/v1.34.1.tar.gz",
        ],
    )

    # # Branch: master
    # # Commit: 99e0cf0e35f21ae9ba601e3c42d133c23abf1c25
    # # Date: 2020-10-15 14:32:58 +0000 UTC
    # # URL: https://github.com/bazelbuild/rules_go/commit/99e0cf0e35f21ae9ba601e3c42d133c23abf1c25
    # #
    # # Announce release 0.24.4, 0.23.12 [skip ci] (#2680)
    # # Size: 516973 (517 kB)
    # http_archive(
    #     name = "io_bazel_rules_go",
    #     sha256 = "7f52bf5679e2d7ae90d25cce25af80707eaac0e497d67f970a0c0704011163db",
    #     strip_prefix = "rules_go-99e0cf0e35f21ae9ba601e3c42d133c23abf1c25",
    #     urls = ["https://github.com/bazelbuild/rules_go/archive/99e0cf0e35f21ae9ba601e3c42d133c23abf1c25.tar.gz"],
    # )

    # Release: v0.24.11
    # TargetCommitish: release-0.24
    # Date: 2021-01-19 23:11:54 +0000 UTC
    # URL: https://github.com/bazelbuild/rules_go/releases/tag/v0.24.11
    # Size: 523890 (524 kB)
    http_archive(
        name = "io_bazel_rules_go",
        sha256 = "d2b5751d8ae55ac011540453cf9da49ee12b832d0a98ca8ffae99285abb481f7",
        strip_prefix = "rules_go-0.24.11",
        urls = ["https://github.com/bazelbuild/rules_go/archive/v0.24.11.tar.gz"],
    )

    http_archive(
        name = "bazel_gazelle",
        sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
        ],
    )

    # Branch: master
    # Commit: a4a1ccffc666db5376342789ad021a943fb84256
    # Date: 2021-01-10 12:54:33 +0000 UTC
    # URL: https://github.com/bazelbuild/rules_python/commit/a4a1ccffc666db5376342789ad021a943fb84256
    #
    # Remove the rules_python_external README (#391)
    # Size: 2563399 (2.6 MB)
    http_archive(
        name = "rules_python",
        sha256 = "b228318a786d99b665bc83bd6cdb81512cae5f8eb15e8cd19f9956604b8939f5",
        strip_prefix = "rules_python-a4a1ccffc666db5376342789ad021a943fb84256",
        urls = ["https://github.com/bazelbuild/rules_python/archive/a4a1ccffc666db5376342789ad021a943fb84256.tar.gz"],
    )

    # Branch: master
    # Commit: 628462deaaf5ea32ebdad9c55ce1bd4c67e9821e
    # Date: 2020-10-27 14:20:56 +0000 UTC
    # URL: https://github.com/rules-proto-grpc/rules_proto_grpc/commit/628462deaaf5ea32ebdad9c55ce1bd4c67e9821e
    #
    # Add CHANGELOG.md
    # Size: 269963 (270 kB)
    http_archive(
        name = "rules_proto_grpc",
        sha256 = "f23f728f4a29a3e60233b0ec9cb28de59ae97dbe407f3067a7b6015e9bd83f7e",
        strip_prefix = "rules_proto_grpc-628462deaaf5ea32ebdad9c55ce1bd4c67e9821e",
        urls = [
            "https://github.com/rules-proto-grpc/rules_proto_grpc/archive/628462deaaf5ea32ebdad9c55ce1bd4c67e9821e.tar.gz",
        ],
    )

    # Release: 0.14.0
    # TargetCommitish: master
    # Date: 2020-08-10 21:14:09 +0000 UTC
    # URL: https://github.com/bazelbuild/rules_swift/releases/tag/0.14.0
    # Size: 149699 (150 kB)
    http_archive(
        name = "build_bazel_rules_swift",
        sha256 = "fa746a50f442ea4bcce78b747182107b4f0041f868b285714364ce4508d19979",
        strip_prefix = "rules_swift-0.14.0",
        urls = [
            "https://github.com/bazelbuild/rules_swift/archive/0.14.0.tar.gz",
        ],
    )
