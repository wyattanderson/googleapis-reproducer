# An attempt at using Bazel, Protobuf, and Go without interference

This is an attempt at getting Bazel, Protobuf, and Go to work nicely together
without too many hacks or workarounds.

## Background

The [`google-cloud-go`
migration](https://github.com/googleapis/google-cloud-go/blob/main/migration.md)
feels like the trigger for all of this. Some packages are migrated, some
aren't, and there are aliases all over the place that occasionally seem to
conflict in strange ways.

With the latest, unpatched versions of Gazelle and `rules_go`, the projects in
this repository are not buildable. They build with native Go tooling, but not
with Bazel.

## The Example projects

- `visionv1/` references the valid-but-not-explicitly-deprecated v1 Google
  Vision API. This target is set up to ensure references to our own Protobuf
  generated code, the `cloud.google.com` package, _and_ the
  `google.golang.org/genproto` package to ensure they all build and link
  correctly.
- `visionv2/` references the v2 Google Vision API as well as some well-known
  types for fun.
- `legacy` pulls in the v1 Go Protobuf well-known types (which are now just
  aliases to the v2 generated code), as well as the v2 `anypb` WKT and our own
  generated Protobuf code.

To reproduce the problems with these projects, rip out the `patches` in
`WORKSPACE`, the directives in `BUILD`, re-run Gazelle, and observe that we
encounter lovely problems like this:

```
external/com_google_cloud_go_longrunning/longrunning.go:60:10: cannot use inner (variable of type *"cloud.google.com/go/longrunning/autogen".OperationsClient) as type operationsClient in struct literal:
        *"cloud.google.com/go/longrunning/autogen".OperationsClient does not implement operationsClient (wrong type for CancelOperation method)
                have CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error
                want CancelOperation(context.Context, *"google.golang.org/genproto/googleapis/longrunning".CancelOperationRequest, ...gax.CallOption) error
```


## Goals

I'm simply not interested in having the build tools apply special handling to
Google's APIs, Protobuf imports, Protobuf well-known types, or anything of the
sort. It's convenient when it works, but when it doesn't work, it's incredibly
frustrating and impossible to remove or disable without patching in Bazel.
This repository intends to illustrate the necessary patches to remove or
disable some of this magic functionality.

I'm sure there are historical reasons for why the special handling is in place
for well-known types and Google APIs, but I don't believe it's necessary
anymore and it conflicts with the native Go tooling. When my Go code imports
a package like `google.golang.org/protobuf/types/known/anypb`, I want Bazel to
follow the same path for resolving that package that the native Go tooling
would follow.

I'm okay with needing to manually specify resolution rules for Protobuf
imports of files like `google/protobuf/any.proto` via Gazelle directives, as
illustrated in `BUILD`. I'd prefer the explicit control of how those imports
resolve to dependencies for `proto_library` and `go_proto_library` targets.
I'm aware that the published Go packages may not exactly match the generated
code from the Protobuf that may be referenced, but I'm fine with this
theoretical limitation because the overwhelming majority of the Go ecosystem
faces the same fate.

## Changes

- I've configured the `//:gazelle-update-repos` target to set
  `build_file_proto_mode = "disable_global"` for all Go dependencies. I want
  the same dependency behavior that I would get with native Go tooling.
- I've patched Gazelle to remove any special handling for Protobuf well-known
  types.
- I've patched Gazelle to remove any special handling for imports of Google
  Cloud Protobuf packages, or any reference to `@go_googleapis`.
- I've patched `rules_go` to remove the `@com_github_golang_protobuf` and
  `@org_golang_google_genproto` repositories and their patches. I want to use
  these repositories directly as a `go_repository`.
- I've patched `rules_go` to change the well-known type targets to _aliases_
  so that third-party packages like [`grpc-gateway` that reference
  `@io_bazel_rules_go//proto/wkt`
  targets](https://github.com/grpc-ecosystem/grpc-gateway/blob/0eb17c3d70415c44406b0f02265ceaad570d6fa3/runtime/BUILD.bazel#L30)
  behave correctly when referenced as a `go_repository`.
- I've included `@com_google_googleapis` directly and specified Gazelle
  directives to point to it as necessary from Protobuf. I've also patched it
  [to fix an incorrect `go_package`
  option](https://github.com/googleapis/googleapis/pull/782).

## The Result

Everything builds under Bazel and under Go native tooling. Rejoice. I think it
would be nice if the ecosystem could move in this direction if there aren't
any major downsides I'm missing. It could be helpful to provide a default set
of directives people can use for setting up Gazelle to understand Go
resolution, but it would be nice if this was explicit instead of magical like
it is now.
