load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def base_deps():
    # Branch: master
    # Commit: 99e0cf0e35f21ae9ba601e3c42d133c23abf1c25
    # Date: 2020-10-15 14:32:58 +0000 UTC
    # URL: https://github.com/bazelbuild/rules_go/commit/99e0cf0e35f21ae9ba601e3c42d133c23abf1c25
    #
    # Announce release 0.24.4, 0.23.12 [skip ci] (#2680)
    # Size: 516973 (517 kB)
    http_archive(
        name = "io_bazel_rules_go",
        sha256 = "7f52bf5679e2d7ae90d25cce25af80707eaac0e497d67f970a0c0704011163db",
        strip_prefix = "rules_go-99e0cf0e35f21ae9ba601e3c42d133c23abf1c25",
        urls = ["https://github.com/bazelbuild/rules_go/archive/99e0cf0e35f21ae9ba601e3c42d133c23abf1c25.tar.gz"],
    )

    http_archive(
        name = "bazel_gazelle",
        sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
        ],
    )

    http_archive(
        name = "com_google_protobuf",
        sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
        strip_prefix = "protobuf-3.11.4",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
    )
