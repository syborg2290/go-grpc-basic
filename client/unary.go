package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func callsSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
