package main

import (
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"log"
	"net"
)

const (
	port = ":4000"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	// Create a TCP listener on port 8080
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the helloServer as the implementation for GreetServiceServer
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	// Print the server address
	log.Printf("Server started at %v", lis.Addr())

	// Start the gRPC server and bind it to the TCP listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
