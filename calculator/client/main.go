package main

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/grpc-golang-examples/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)
	DoSum(client)
}

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
