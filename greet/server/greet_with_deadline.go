package main

import (
	"context"
	"log"
	"time"

	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked wit %v\n", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request!")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Helllo " + in.FirstName,
	}, nil
}
