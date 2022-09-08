package main

import (
	"log"
	"net"

	pb "github.com/luizmoitinho/grpc-golang-examples/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
	how to use evans: https://github.com/ktr0731/evans#macos
		- evans --host localhost --port 50051 --reflection repl
		- show services
		- service CalculatorService
		- call Sum
		(controld+d done operation)
*/

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}

}
