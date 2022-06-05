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

func (s *Server) Primes(in *pb.CalculatorPrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	var k int32 = 2
	number := in.Number
	for number > 1 {
		if number%k == 0 {
			log.Printf("%d ", k)
			stream.Send(&pb.CalculatorResponse{
				Result: k,
			})
			number = number / k
		} else {
			k++
		}
	}
	return nil
}
