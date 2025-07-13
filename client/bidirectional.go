package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	stream, err := client.SayHelloBidirectional(context.Background())
	if err != nil {
		log.Fatalf("Error creating bidirectional stream: %v", err)
	}

	done := make(chan struct{})

	// Send names to the server
	go func() {
		for _, name := range names.Names {
			req := &pb.HelloRequest{Name: name}
			if err := stream.Send(req); err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
			log.Printf("Sent: %s", name)
			time.Sleep(2 * time.Second) // Simulate some processing delay
		}
		stream.CloseSend()
	}()

	// Receive responses from the server
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server finished sending all responses")
				break
			}
			if err != nil {
				log.Fatalf("Error receiving response: %v", err)
			}
			log.Printf("Received message: %s", resp.Message)
		}
		close(done)
	}()

	<-done // Wait for the receiving goroutine to finish
}
