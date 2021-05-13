package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kirtimd/grpc-go/helloworld"
)

const (
	port = ":50051"
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
	savedGreetings []*pb.Greeting
}

func (s *helloWorldServer) Greet(c context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", request.GetName())
	return &pb.GreetResponse{Message: "Hey " + request.GetName()}, nil
}

//server-side streaming
func (s *helloWorldServer) GreetInManyLanguages(request *pb.GreetRequest, stream pb.HelloWorld_GreetInManyLanguagesServer) error {
	for _, greeting := range s.savedGreetings {
		response := &pb.GreetResponse{Message: greeting.Expression + ", " + request.GetName()}
		if error := stream.Send(response); error != nil {
			return error
		}
	}
	return nil
}

//client-side streaming
func (s *helloWorldServer) GreetManyPeople(stream pb.HelloWorld_GreetManyPeopleServer) error {
	log.Print("Client streamed: ")
	message := "Hello "
	for {
		request, error := stream.Recv()

		if error == io.EOF {
			message += "!"
			return stream.SendAndClose(&pb.GreetResponse{Message: message})
		}
		if error != nil {
			return error
		}

		//before every name (except the 1st one), add a comma
		if message != "Hello " {
			message += ", "
		}

		log.Print(request.Name + " ")

		message += request.Name
	}
}

//bidirectional streaming
//both streams are independent
//order is preserved
//client and server can either wait for a stream to be complete
//or read a message, write a message, and so on...
func (s *helloWorldServer) Chat(stream pb.HelloWorld_ChatServer) error {
	for {
		request, error := stream.Recv()
		//Note: request is of type Talk
		if error == io.EOF {
			return nil
		}

		if error != nil {
			return error
		}

		i := rand.Intn(len(s.savedGreetings))

		response := &pb.Talk{Sentence: s.savedGreetings[i].Expression + ", " + request.Sentence}
		if error := stream.Send(response); error != nil {
			return error
		}

	}
}

func newServer() *helloWorldServer {

	s := &helloWorldServer{}
	s.savedGreetings = []*pb.Greeting{
		&pb.Greeting{Expression: "Namaskar"},
		&pb.Greeting{Expression: "Howdy"},
		&pb.Greeting{Expression: "Ola"},
		&pb.Greeting{Expression: "Olleya Dina"},
		&pb.Greeting{Expression: "Guten Tag"},
	}
	return s
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, newServer())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
