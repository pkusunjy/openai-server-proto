load('@rules_proto//proto:defs.bzl', 'proto_library')
load('@com_github_grpc_grpc//bazel:cc_grpc_library.bzl', 'cc_grpc_library')

package(
    default_visibility = ['//visibility:public'],
)

# auth
proto_library(
    name = 'auth_proto',
    srcs = ['auth.proto'],
)

cc_proto_library(
    name = 'auth_cc_proto',
    deps = [':auth_proto'],
)

cc_grpc_library(
    name = 'auth_cc_grpc',
    srcs = [':auth_proto'],
    grpc_only = True,
    deps = [':auth_cc_proto'],
)

# chat_completion
proto_library(
    name = 'chat_completion_proto',
    srcs = ['chat_completion.proto'],
)

cc_proto_library(
    name = 'chat_completion_cc_proto',
    deps = [':chat_completion_proto'],
)

cc_grpc_library(
    name = 'chat_completion_cc_grpc',
    srcs = [':chat_completion_proto'],
    grpc_only = True,
    deps = [':chat_completion_cc_proto'],
)
