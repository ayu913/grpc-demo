package main

import (
	pb "grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	Port = ":3006"
)

type helloServer struct {
	pb.UnimplementedGreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
