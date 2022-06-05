package main

import (
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
	//DoSum(client)
	//DoPrimes(client)
	//DoAvg(client)
	DoMax(client)
}
