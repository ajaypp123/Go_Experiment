package main

import (
	"log"
	"net"

	pb "github.com/ajaypp/basic-go-grpc/proto"
	"github.com/ajaypp/basic-go-grpc/server/communicate"
	"google.golang.org/grpc"
)

// define the port
const (
	port = ":8080"
)

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &communicate.HelloServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
