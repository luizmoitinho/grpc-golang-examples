package main

import (
	"context"
	"log"
	"time"

	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DoGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("DoGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Luiz",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded")
				return
			} else {
				log.Fatalf("Unexepcted gRPC error: %v", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}
	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
