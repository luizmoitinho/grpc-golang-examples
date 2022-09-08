package main

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
)

func DoGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Luiz Carlos",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
