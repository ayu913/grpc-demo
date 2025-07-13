package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Println("SayHello called")
	return &pb.HelloResponse{Message: "Hello from gRPC server!"}, nil
}


