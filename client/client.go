package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/kirtimd/grpc-go"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	connection, error := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if error != nil {
		log.Fatalf("Did not connect: %v", error)
	}
	defer connection.Close()
	c := pb.NewClient(connection)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, error := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if error != nil {
		log.Fatalf("Could not greet: %v", error)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
