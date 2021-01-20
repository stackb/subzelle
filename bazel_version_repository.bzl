"""bazel_version_repository captures the bazel version as an importable symbol

bazel_version_repository is necessary to satisfy a requirement for
@upb (a dependency of @com_github_grpc_grpc).
"""

def _bazel_version_repository(repository_ctx):
    repository_ctx.file("BUILD", "exports_files(['def.bzl'])")
    repository_ctx.file("bazel_version.bzl", "bazel_version = \"{}\"".format(native.bazel_version))
    repository_ctx.file("def.bzl", "BAZEL_VERSION='{}'".format(native.bazel_version))

bazel_version_repository = repository_rule(
    implementation = _bazel_version_repository,
)
