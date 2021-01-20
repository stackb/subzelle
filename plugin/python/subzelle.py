import grpc

from concurrent import futures

from proto import (
    language_pb2 as lpb,
    language_pb2_grpc as lgrpc,
)

class LanguageServicer(lgrpc.LanguageServicer):
    # 
    def Create(self, request, context):
        return lpb.CreateResponse(id=1001, name=request.name)

def main():
    address = "[::]:50051"
    print("Starting subplugin gRPC server on " + address)
    # Create and start the server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    lgrpc.add_LanguageServicer_to_server(LanguageServicer(), server)
    server.add_insecure_port(address)
    server.start()
    server.wait_for_termination()
    # print("Stopping subplugin gRPC server.")

if __name__ == "__main__":
    main()
