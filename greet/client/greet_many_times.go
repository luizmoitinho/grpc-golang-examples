package main

import (
	"context"
	"io"
	"log"

	"github.com/luizmoitinho/grpc-golang-examples/greet/proto"
)

func DoGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &proto.GreetRequest{
		FirstName: "Luiz",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg)
	}
}
