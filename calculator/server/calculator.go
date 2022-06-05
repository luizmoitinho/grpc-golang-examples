package main

import (
	"context"
	"fmt"
	"io"
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

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	fmt.Println("Avg function was invoked")
	var sumValues float64 = 0
	var counterValues int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&pb.AvgResponse{
					Result: sumValues / float64(counterValues),
				},
			)
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Receiving: %v\n", req)
		sumValues += float64(req.Number)
		counterValues++
	}
}
