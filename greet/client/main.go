package main

import (
	"log"

	pb "github.com/luizmoitinho/grpc-golang-examples/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("error while loading CA trust certificate: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	DoGreet(client)
	//DoGreetManyTimes(client)
	//DoLongGreet(client)
	//DoGreetEveryone(client)

	//DoGreetWithDeadline(client, 5*time.Second)
	//DoGreetWithDeadline(client, 1*time.Second) // should be return a deadline error

}
