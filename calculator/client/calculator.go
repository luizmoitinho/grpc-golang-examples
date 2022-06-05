package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/luizmoitinho/grpc-golang-examples/calculator/proto"
)

func DoSum(c pb.CalculatorServiceClient) {
	log.Println("DoSum was invoked")
	input := &pb.CalculatorRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := c.Sum(context.Background(), input)

	if err != nil {
		log.Fatalf("Could not Sum: %v", err)
	}

	log.Printf("Sum: %d + %d = %d", input.FirstNumber, input.SecondNumber, res.Result)
}

func DoPrimes(c pb.CalculatorServiceClient) {
	log.Println("DoPrimes was invoked")
	req := &pb.CalculatorPrimesRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	log.Printf("Primes of %d: \n", req.Number)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("%d ", msg.Result)
	}

}

func DoAvg(c pb.CalculatorServiceClient) {
	log.Println("DoAvg was invoked")
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Erro while calling Avg: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}
	log.Printf("Avg: %f\n", res.Result)
}

func DoMax(c pb.CalculatorServiceClient) {
	log.Println("DoMax was invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Errpr while creating stream: %v\n", err)
	}

	//1, 5, 3, 6, 2, 20,
	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

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
				log.Printf("Error while receiving: %v\n", res.Result)
				break
			}
			log.Printf("Receveid: %v\n", res.Result)
		}
		close(waitc)

	}()

	<-waitc

}
