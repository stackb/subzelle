import grpc

from concurrent import futures

from proto import (
    language_pb2 as lpb,
    language_pb2_grpc as lgrpc,
)

class LanguageServicer(lgrpc.LanguageServicer):
    def Kinds(self, request, context):
        print("Kinds <-")
        return lpb.KindsResponse(kinds={
            "py_binary": lpb.KindInfo(
                match_any=True,
                non_empty_attrs={"srcs": True, "deps": True},
                mergeable_attrs={"srcs": True},
                resolve_attrs={"deps": True},
            ),
            "py_library": lpb.KindInfo(
                match_any=True,
                non_empty_attrs={"srcs": True, "deps": True},
                mergeable_attrs={"srcs": True},
                resolve_attrs={"deps": True},
            ),
        })

    def Loads(self, request, context):
        print("Loads <-")
        return lpb.LoadsResponse(Load=[
            lpb.LoadInfo(
                name="@rules_python//python:defs.bzl",
                symbols=["py_binary", "py_library", "py_test"],
            ),
        ])

    def RegisterFlags(self, request, context):
        print("RegisterFlags <-")
        return request.config_flag_set

    def CheckFlags(self, request, context):
        print("CheckFlags <-")
        return request

    def KnownDirectives(self, request, context):
        print("KnownDirectives <-")
        return lpb.KnownDirectivesResponse(directive=[])

    def Configure(self, request, context):
        print("Configure <-")
        return request

    def GenerateRules(self, request, context):
        print("GenerateRules <- %r" % request)
        return lpb.GenerateResult()


def main():
    address = "[::]:50051"
    print("Starting subplugin gRPC server on " + address)
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    lgrpc.add_LanguageServicer_to_server(LanguageServicer(), server)
    server.add_insecure_port(address)
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    main()
