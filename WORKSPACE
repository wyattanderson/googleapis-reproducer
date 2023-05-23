load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    patch_args = [
        "-E",
        "-p1",
    ],
    patches = [
        "//patches:io_bazel_rules_go-remove-golang-protobuf.patch",
        "//patches:io_bazel_rules_go-alias.patch",
    ],
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    patch_args = [
        "-E",
        "-p1",
    ],
    patches = ["//patches:bazel_gazelle-ignore-known-imports.patch"],
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

http_archive(
    name = "aspect_bazel_lib",
    sha256 = "f8fa3193009232ca989de21964ea860c8b3279ec73ba6eff456d8bf61fb3ab1f",
    strip_prefix = "bazel-lib-1.27.0",
    url = "https://github.com/aspect-build/bazel-lib/releases/download/v1.27.0/bazel-lib-v1.27.0.tar.gz",
)

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")
load("//:deps.bzl", "go_deps")

# gazelle:repository_macro deps.bzl%go_deps

aspect_bazel_lib_dependencies()

go_deps()

go_rules_dependencies()

go_register_toolchains(version = "1.19.5")

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_google_googleapis",
    sha256 = "9a106bbb387d2d6eeb7bcd5a4494cff6fa892eb9e92f128d9daf9f0515392f98",
    strip_prefix = "googleapis-117be9dfdf65ff766a794c8b85d5d7480a1fd83d",
    urls = [
        "https://github.com/googleapis/googleapis/archive/117be9dfdf65ff766a794c8b85d5d7480a1fd83d.zip",
    ],
)

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
)
