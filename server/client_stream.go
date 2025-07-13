package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

// SayHelloClientStreaming implements the client streaming RPC for the GreetService.
// It receives a stream of GreetRequest messages from the client, appends a greeting
// for each received name to a list, and once the client has finished sending messages,
// it responds with a MessagesList containing all the greetings. If an error occurs
// while receiving messages, it logs the error and returns it.

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{
				Messages: messages,
			})
		}
		if err != nil {
			log.Printf("Error receiving client stream: %v", err)
			return err
		}

		log.Printf("Received: %s", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}


// stream.Recv() is used to receive messages from the client stream.
// stream.SendAndClose() is used to send a final response back to the client after all messages have been received.