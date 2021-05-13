package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kirtimd/grpc-go/helloworld"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func greetOnce(client pb.HelloWorldClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, error := client.Greet(ctx, &pb.GreetRequest{Name: name})
	if error != nil {
		log.Fatalf("Could not greet: %v", error)
	}
	log.Printf("Greeting: %s", response.GetMessage())

}

//client sends a stream to server
func greetManyPeopleOnce(client pb.HelloWorldClient, names []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, error := client.GreetManyPeople(ctx)
	if error != nil {
		log.Fatalf("GreetManyPeople Error: %v", error)
	}
	for _, n := range names {
		request := &pb.GreetRequest{Name: n}
		if error := stream.Send(request); error != nil {
			log.Fatalf("greetManyPeople Error: %v", error)
		}
	}
	reply, error := stream.CloseAndRecv()
	if error != nil {
		log.Fatalf("CloseAndRecv Error: %v", error)
	}
	log.Printf("%v", reply)
}

//client receives a stream from server
func greetInManyLanguages(client pb.HelloWorldClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.GreetRequest{Name: name}
	stream, error := client.GreetInManyLanguages(ctx, request)
	if error != nil {
		log.Fatalf("GreetInManyLanguages Error: %v", error)
	}

	for {
		response, error := stream.Recv()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatalf("stream.Recv Error: %v", error)
		}
		log.Println(response)
	}
}

func chat(client pb.HelloWorldClient, names []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, error := client.Chat(ctx)
	if error != nil {
		log.Fatalf("Char error: %v", error)
	}
	//channel for sending struct values to the another thread
	ch := make(chan struct{})

	go func() { //starts another thread of execution for the code enclosed
		for {
			response, error := stream.Recv()
			if error == io.EOF {
				close(ch) //close channel
				return
			}
			if error != nil {
				log.Fatalf("Recv failed: %v", error)
			}

			log.Printf("Received message: %v", response.Sentence)
		}
	}()

	for _, name := range names {
		if error := stream.Send(&pb.Talk{Sentence: name}); error != nil {
			log.Fatalf("Send failed: %v", error)
		}
	}
	stream.CloseSend()
	<-ch //blocks this thread until the anonymous thread above finishes
}

func main() {
	connection, error := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if error != nil {
		log.Fatalf("Did not connect: %v", error)
	}
	defer connection.Close()
	client := pb.NewHelloWorldClient(connection)

	names := []string{"Karakoram", "Kanchenjunga", "Garhwal", "Siachen", "Kumaon", "Saltoro"}

	log.Println("-- Single request/response --")
	greetOnce(client, names[0])
	log.Println("-- Client-side streaming --")
	greetManyPeopleOnce(client, names)
	log.Println("-- Server-side streaming --")
	greetInManyLanguages(client, names[1])
	log.Println("-- Bidirectional streaming --")
	chat(client, names)
}
