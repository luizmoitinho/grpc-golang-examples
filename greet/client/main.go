package main

import (
	"log"
	"time"

	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
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

	client := pb.NewGreetServiceClient(conn)
	//DoGreet(client)
	//DoGreetManyTimes(client)
	//DoLongGreet(client)
	//DoGreetEveryone(client)

	//DoGreetWithDeadline(client, 5*time.Second)
	DoGreetWithDeadline(client, 1*time.Second) // should be return a deadline error

}
