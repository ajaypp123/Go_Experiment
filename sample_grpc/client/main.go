package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ajaypp123/sample_grpc/client/communication"
	pb "github.com/ajaypp123/sample_grpc/proto/service"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	// Create a gRPC channel to communicate with the server.
	// Replace "localhost:50051" with the address of your server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Create a client stub for the Communication service.
	empClient := communication.EmployeeClient{
		Client: pb.NewEmployeeServiceClient(conn),
	}
	// Call all operations
	empClient.GetEmployee(ctx)
	empClient.GetAllEmployee(ctx)
	empClient.AddEmployee(ctx)
	empClient.Chat(ctx)

	// Create a client stub for the Communication service.
	studClient := communication.StudentClient{
		Client: pb.NewStudentServiceClient(conn),
	}
	// Call all operations
	studClient.GetStudent(ctx)
	studClient.GetAllStudent(ctx)
	studClient.AddStudent(ctx)
	studClient.Chat(ctx)

	// =======================================
	// ============= Async Single Response ===================
	// =======================================
	// create a channel for the response
	respChan := make(chan *pb.StudentResponse, 1)
	go studClient.AsyncGetStudent(ctx, respChan)
	fmt.Println("************** After Work AsyncGetStudent *************")
	resp := <-respChan
	fmt.Println("Response received:", resp)

	// =======================================
	// ============= Async Multiple Response ===================
	// =======================================
	// create a channel for the response
	respsChan := make(chan *pb.StudentResponse)
	go studClient.AsyncGetAllStudent(ctx, respsChan)
	fmt.Println("************** After Work AsyncGetAllStudent *************")
	// Wait for responses and print them
	for response := range respsChan {
		log.Printf("StudentResponse: %v", response)
	}
}
