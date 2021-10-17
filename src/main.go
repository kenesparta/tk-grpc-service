package main

import (
	"github.com/kenesparta/tkGrpcService/gen/multiply"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcPort = ":8085"

func main() {
	lis, err:= net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s - %v\n", grpcPort, err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	multiply.RegisterMultiplyServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC over port %s - %v\n", grpcPort, err)
	}
}
