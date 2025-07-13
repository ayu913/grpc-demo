package main

import (
	pb "grpc-demo/proto"
	"time"
)

// SayHelloServerStreaming implements the server-side streaming RPC for greeting multiple names.
// It receives a NamesList request containing a list of names, and for each name, sends a HelloResponse
// message back to the client via the stream. Each response contains a personalized greeting message.
// A 2-second delay is introduced between sending each message to simulate processing time.
// Returns an error if sending a message on the stream fails.

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	for _, name := range req.Names {
		resp := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // Simulate some processing delay
	}
	return nil
}

// stream.Send() is used to send messages back to the client stream.
