package main

import (
	"log"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":3006"

func main() {

	clientConn, err := grpc.NewClient("localhost:3006", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer clientConn.Close()

	client := pb.NewGreetServiceClient(clientConn)

	names := &pb.NamesList{
		Names: []string{"Alice", "Bob", "Charlie"},
	}

	// CallSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
