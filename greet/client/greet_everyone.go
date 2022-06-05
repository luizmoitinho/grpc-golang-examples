package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/luizmoitinho/grpc-golang-examples/greet/proto"
)

func DoGreetEveryone(c proto.GreetServiceClient) {
	log.Println("DoGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Errpr while creating stream: %v\n", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Luiz"},
		{FirstName: "João"},
		{FirstName: "Eduardo"},
		{FirstName: "Henrique"},
	}

	waitc := make(chan struct{})

	//send stream of requests
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
