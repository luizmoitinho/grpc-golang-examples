package main

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	DoGreet(client)
}

func DoGreet(c pb.GreetServiceClient) {
	log.Println("doGreet  was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Luiz Carlos",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
