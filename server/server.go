package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kirtimd/grpc-go/helloworld"
)

const (
	port = ":50051"
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
}

func (s *helloWorldServer) SayHello(c context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", request.GetName())
	return &pb.HelloReply{Message: "Hey " + request.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloWorldServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
