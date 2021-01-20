# subzelle

An experimental gazelle plugin that adapts the gazelle-plugin interface over
gRPC.

Ever wanted to write a bazel `BUILD` file generator?
[Gazelle](https://github.com/bazelbuild/bazel-gazelle) is a great tool to do
that. Gazelle implements a plugin architecture, but until now, plugins always had to be written in go.

With `subzelle`, you write your gazelle plugin in the language of your choice in
the form of a gRPC server. The subzelle plugin that is statically compiled into
gazelle will then RPC out to your plugin.

### Example

Currently the example requires 2 terminals. The final version will support launching your subprocess.

```
$ bazel run //plugin/python:subzelle
Starting subplugin gRPC server on [::]:50051
```

```
$ bzl run example/python:subzelle --build_python_zip
```

Client Output (gazelle/subzelle):

```
gazelle: Kinds ->
gazelle: Loads ->
gazelle: RegisterFlags ->
gazelle: CheckFlags ->
gazelle: KnownDirectives ->
gazelle: /home/pcj/go/src/github.com/stackb/subzelle/BUILD.bazel: unknown directive: gazelle:prefix
gazelle: Configure ->
gazelle: Configure ->
gazelle: GenerateRules ->
gazelle: Configure ->
gazelle: Configure ->
gazelle: GenerateRules ->
gazelle: GenerateRules ->
gazelle: Configure ->
gazelle: Configure ->
gazelle: GenerateRules ->
gazelle: GenerateRules ->
gazelle: Configure ->
gazelle: Configure ->
gazelle: GenerateRules ->
gazelle: GenerateRules ->
gazelle: Configure ->
gazelle: /home/pcj/go/src/github.com/stackb/subzelle/plugin/python/BUILD.bazel: unknown directive: gazelle:proto
gazelle: Configure ->
gazelle: GenerateRules ->
gazelle: GenerateRules ->
gazelle: Configure ->
gazelle: GenerateRules ->
gazelle: GenerateRules ->
```

Server Output (python):

```
Kinds <-
Loads <-
RegisterFlags <-
CheckFlags <-
KnownDirectives <-
Configure <-
Configure <-
GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/.vscode"
rel: ".vscode"
regular_files: "settings.json"

Configure <-
Configure <-
GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/cmd/subzelle"
rel: "cmd/subzelle"
file {
  pkg: "cmd/subzelle"
  path: "/home/pcj/go/src/github.com/stackb/subzelle/cmd/subzelle/BUILD.bazel"
  load {
    name: "@bazel_gazelle//:def.bzl"
  }
  rule {
    kind: "gazelle_binary"
    name: "subzelle"
  }
  rule {
    kind: "gazelle"
    name: "gazelle"
  }
  content: "load(\"@bazel_gazelle//:def.bzl\", \"gazelle\", \"gazelle_binary\")\n\ngazelle_binary(\n    name = \"subzelle\",\n    # NOTE: DEFAULT_LANGUAGES intentionally excluded here - this won\'t do go or proto.\n    languages = [\"//pkg/subzelle:go_default_library\"],\n    visibility = [\"//visibility:public\"],\n)\n\ngazelle(\n    name = \"gazelle\",\n    gazelle = \":subzelle\",\n    prefix = \"github.com/stackb/subzelle\",\n)\n"
}
regular_files: "BUILD.bazel"

GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/cmd"
rel: "cmd"
subdirs: "subzelle"

Configure <-
Configure <-
GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/example/python"
rel: "example/python"
file {
  pkg: "example/python"
  path: "/home/pcj/go/src/github.com/stackb/subzelle/example/python/BUILD.bazel"
  load {
    name: "//:def.bzl"
  }
  rule {
    kind: "subzelle"
    name: "subzelle"
  }
  content: "load(\"//:def.bzl\", \"subzelle\")\n\nsubzelle(\n    name = \"subzelle\",\n    plugin = \"//plugin/python\",\n    prefix = \"github.com/stackb/subzelle/example/python\",\n)\n"
}
regular_files: "BUILD.bazel"

GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/example"
rel: "example"
subdirs: "python"

Configure <-
Configure <-
GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/pkg/subzelle"
rel: "pkg/subzelle"
file {
  pkg: "pkg/subzelle"
  path: "/home/pcj/go/src/github.com/stackb/subzelle/pkg/subzelle/BUILD.bazel"
  load {
    name: "@io_bazel_rules_go//go:def.bzl"
  }
  load {
    name: "@io_bazel_rules_go//proto:def.bzl"
  }
  rule {
    kind: "go_library"
    name: "go_default_library"
  }
  rule {
    kind: "go_proto_library"
    name: "language_go_proto"
  }
  content: "load(\"@io_bazel_rules_go//go:def.bzl\", \"go_library\")\nload(\"@io_bazel_rules_go//proto:def.bzl\", \"go_proto_library\")\n\ngo_library(\n    name = \"go_default_library\",\n    srcs = [\n        \"configurer.go\",\n        \"fix.go\",\n        \"generator.go\",\n        \"language.go\",\n        \"plugin.go\",\n        \"process.go\",\n        \"resolver.go\",\n        \"types.go\",\n    ],\n    importpath = \"github.com/stackb/subzelle/pkg/subzelle\",\n    visibility = [\"//visibility:public\"],\n    deps = [\n        \":language_go_proto\",\n        \"@bazel_gazelle//config:go_default_library\",\n        \"@bazel_gazelle//label:go_default_library\",\n        \"@bazel_gazelle//language:go_default_library\",\n        \"@bazel_gazelle//repo:go_default_library\",\n        \"@bazel_gazelle//resolve:go_default_library\",\n        \"@bazel_gazelle//rule:go_default_library\",\n        \"@com_github_golang_protobuf//proto:go_default_library\",\n        \"@org_golang_google_grpc//:go_default_library\",\n    ],\n)\n\ngo_proto_library(\n    name = \"language_go_proto\",\n    compilers = [\"@io_bazel_rules_go//proto:go_grpc\"],\n    importpath = \"github.com/stackb/subzelle/language\",\n    proto = \"//proto:language_proto\",\n    visibility = [\"//visibility:public\"],\n)\n"
}
regular_files: "BUILD.bazel"
regular_files: "cheetah.go"
regular_files: "config.go"
regular_files: "configurer.go"
regular_files: "fix.go"
regular_files: "generator.go"
regular_files: "kind.go"
regular_files: "language.go"
regular_files: "plugin.go"
regular_files: "process.go"
regular_files: "resolver.go"
regular_files: "types.go"

GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/pkg"
rel: "pkg"
file {
  pkg: "pkg"
  path: "/home/pcj/go/src/github.com/stackb/subzelle/pkg/BUILD.bazel"
}
subdirs: "subzelle"
regular_files: "BUILD.bazel"

Configure <-
Configure <-
GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/plugin/python"
rel: "plugin/python"
file {
  pkg: "plugin/python"
  path: "/home/pcj/go/src/github.com/stackb/subzelle/plugin/python/BUILD.bazel"
  directive {
    key: "proto"
    value: "disable"
  }
  load {
    name: "@rules_python//python:defs.bzl"
  }
  load {
    name: "//:plugin.bzl"
  }
  load {
    name: "@rules_proto_grpc//python:defs.bzl"
  }
  rule {
    kind: "subzelle_plugin"
    name: "python"
  }
  rule {
    kind: "py_binary"
    name: "subzelle"
  }
  rule {
    kind: "python_grpc_library"
    name: "language_python_grpc"
  }
  content: "load(\"@rules_python//python:defs.bzl\", \"py_binary\")\nload(\"//:plugin.bzl\", \"subzelle_plugin\")\nload(\"@rules_proto_grpc//python:defs.bzl\", \"python_grpc_library\")\n\nsubzelle_plugin(\n    name = \"python\",\n    executable = \":subzelle\",\n    address = \"0.0.0.0:50051\",\n    visibility = [\"//visibility:public\"],\n)\n\npy_binary(\n    name = \"subzelle\",\n    srcs = [\"subzelle.py\"],\n    deps = [\":language_python_grpc\"],\n)\n\n# gazelle:proto disable\n\n# py_test(\n#     name = \"python_grpc_example_test\",\n#     srcs = [\"python_grpc_example_test.py\"],\n#     deps = [\":simple_service_python_grpc_library\"],\n# )\n\npython_grpc_library(\n    name = \"language_python_grpc\",\n    deps = [\"//proto:language_proto\"],\n)\n"
}
regular_files: "BUILD.bazel"
regular_files: "subzelle.py"

GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/plugin"
rel: "plugin"
subdirs: "python"

Configure <-
GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle/proto"
rel: "proto"
file {
  pkg: "proto"
  path: "/home/pcj/go/src/github.com/stackb/subzelle/proto/BUILD.bazel"
  load {
    name: "@rules_proto//proto:defs.bzl"
  }
  rule {
    kind: "proto_library"
    name: "language_proto"
  }
  content: "load(\"@rules_proto//proto:defs.bzl\", \"proto_library\")\n\nproto_library(\n    name = \"language_proto\",\n    srcs = [\"language.proto\"],\n    visibility = [\"//visibility:public\"],\n    deps = [\"@com_google_protobuf//:struct_proto\"],\n)\n"
}
regular_files: "BUILD.bazel"
regular_files: "language.proto"

GenerateRules <- config {
  repo_root: "/home/pcj/go/src/github.com/stackb/subzelle"
  repo_name: "build_stack_subzelle"
  valid_build_file_names: "BUILD.bazel"
  valid_build_file_names: "BUILD"
  index_libraries: true
}
dir: "/home/pcj/go/src/github.com/stackb/subzelle"
file {
  path: "/home/pcj/go/src/github.com/stackb/subzelle/BUILD.bazel"
  directive {
    key: "prefix"
    value: "github.com/stackb/subzelle"
  }
  load {
    name: "@bazel_gazelle//:def.bzl"
  }
  rule {
    kind: "gazelle"
    name: "gazelle"
  }
  rule {
    kind: "exports_files"
  }
  content: "load(\"@bazel_gazelle//:def.bzl\", \"gazelle\")\n\n# gazelle:prefix github.com/stackb/subzelle\ngazelle(\n    name = \"gazelle\",\n    prefix = \"github.com/stackb/subzelle\",\n)\n\nexports_files([\"subzelle.bash.in\"])\n"
}
subdirs: ".vscode"
subdirs: "cmd"
subdirs: "example"
subdirs: "pkg"
subdirs: "plugin"
subdirs: "proto"
regular_files: ".bazelrc"
regular_files: ".bazelversion"
regular_files: "BUILD.bazel"
regular_files: "WORKSPACE"
regular_files: "bazel-bin"
regular_files: "bazel-out"
regular_files: "bazel-subzelle"
regular_files: "bazel-testlogs"
regular_files: "bazel_version_repository.bzl"
regular_files: "def.bzl"
regular_files: "deps.bzl"
regular_files: "go_deps.bzl"
regular_files: "launch.bazelrc"
regular_files: "plugin.bzl"
regular_files: "subzelle.bash.in"
```
