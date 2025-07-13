package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (s *helloServer) SayHelloBidirectional(stream grpc.BidiStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server finished sending all responses")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}
		log.Println("Got requedst with name:", req.Name)

		resp := &pb.HelloResponse{Message: "Hello " + req.Name}
		if err := stream.Send(resp); err != nil {
			log.Fatalf("Error sending response: %v", err)
		}
		log.Printf("Sent response: %s", resp.Message)

	}

	return nil
}
