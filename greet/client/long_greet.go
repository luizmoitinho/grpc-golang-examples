package main

import (
	"context"
	"log"
	"time"

	"github.com/luizmoitinho/grpc-golang-examples/greet/proto"
	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
)

func DoLongGreet(c proto.GreetServiceClient) {
	log.Println("DoLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Luiz"},
		{FirstName: "João"},
		{FirstName: "Eduardo"},
		{FirstName: "Henrique"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Erro while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: \n%s\n", res.Result)
}
