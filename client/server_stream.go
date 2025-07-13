package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server finished sending all responses")
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}
		log.Printf("Greeting: %s", resp.Message)
	}
}
