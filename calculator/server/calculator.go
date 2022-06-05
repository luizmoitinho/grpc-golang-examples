package main

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/grpc-golang-examples/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Sum function was invoked with %v", in)
	return &pb.CalculatorResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
