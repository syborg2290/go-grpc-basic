package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/proto"
	"log"
)

const (
	port = ":4000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"Kasun", "Alice", "Bob"},
	// }

	callsSayHello(client)
}
