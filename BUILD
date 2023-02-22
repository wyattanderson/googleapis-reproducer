load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/wyattanderson/googleapis-reproducer
gazelle(name = "gazelle")

go_library(
    name = "googleapis-reproducer_lib",
    srcs = ["main.go"],
    importpath = "github.com/wyattanderson/googleapis-reproducer",
    visibility = ["//visibility:private"],
    deps = [
        "@com_google_cloud_go_vision//apiv1",
        "@go_googleapis//google/cloud/vision/v1:vision_go_proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "googleapis-reproducer",
    embed = [":googleapis-reproducer_lib"],
    visibility = ["//visibility:public"],
)
