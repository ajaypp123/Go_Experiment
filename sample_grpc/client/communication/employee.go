package communication

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ajaypp123/sample_grpc/proto/service"
)

type EmployeeClient struct {
	Client pb.EmployeeServiceClient
}

// ######################### Unary ##########################
// GetEmployee Call the Unary RPC method. Get Single employee.
func (e *EmployeeClient) GetEmployee(ctx context.Context) {
	fmt.Println()
	fmt.Println("============== Unary ===================")
	fmt.Println("Entry.........")
	request := &pb.GetEmployeeRequest{EmployeeId: "1"}
	response, err := e.Client.GetEmployee(ctx, request)
	if err != nil {
		log.Fatalf("Failed to call Unary: %v", err)
	}
	fmt.Printf("Unary Request: %s\n", request)
	fmt.Printf("Unary Response: %s\n", response)
	fmt.Println("Exit.........")
}

// ######################### Server Streaming ##########################
// Call the ServerStreaming RPC method. GetMultiple employee.
func (e *EmployeeClient) GetAllEmployee(ctx context.Context) {
	fmt.Println()
	fmt.Println("============== Server Streaming ===================")
	fmt.Println("Entry.........")
	serverStreamingRequest := &pb.GetAllEmployeeRequest{
		EmployeeIds: []string{"1", "2", "3"},
	}
	fmt.Printf("ServerStreaming Request: %s\n", serverStreamingRequest)
	serverStreamingStream, err := e.Client.GetAllEmployees(ctx, serverStreamingRequest)
	if err != nil {
		log.Fatalf("Failed to call ServerStreaming: %v", err)
	}
	// var emps []string
	for {
		response, err := serverStreamingStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive ServerStreaming response: %v", err)
		}
		fmt.Printf("ServerStreaming Response: %s\n", response)
		// emps = append(emps, response.Employee.EmployeeId)
		/*
			if len(emps) == 3 {
				break
			}
		*/
	}
	fmt.Println("Exit.........")
}

// ######################### Client Streaming ##########################
// Call the ClientStreaming RPC method. Post multiple employee.
func (e *EmployeeClient) AddEmployee(ctx context.Context) {
	fmt.Println()
	fmt.Println("============== Client Streaming ===================")
	fmt.Println("Entry.........")
	stream, err := e.Client.AddEmployees(ctx)
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}

	// Send some requests.
	for i := 6; i < 9; i++ {
		req := &pb.AddEmployeeRequest{
			Employee: &pb.Employee{
				EmployeeId: fmt.Sprintf("%d", i),
				Name:       fmt.Sprintf("Employee %d", i),
				Salary:     float32(1000 + i*500),
				Company:    "MyCompany",
			},
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
		fmt.Printf("ClientStreaming Request: %s\n", req)
	}

	// Close the stream and receive the response.
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}
	log.Printf("ClientStreaming Response: %v", resp)
	fmt.Println("Exit.........")
}

// ######################### Bidirectional Streaming ##########################
// Call the BidirectionalStreaming RPC method.
func (e *EmployeeClient) Chat(ctx context.Context) {
	fmt.Println()
	fmt.Println("============== Bidirectional Streaming ===================")
	fmt.Println("Entry.........")
	bidirectionalstream, err := e.Client.Chat(ctx)
	if err != nil {
		log.Fatalf("Failed to call BidirectionalStreaming: %v", err)
	}
	for i := 1; i < 15; i++ {
		request := &pb.GetEmployeeRequest{EmployeeId: fmt.Sprintf("%d", i)}
		err := bidirectionalstream.Send(request)
		if err != nil {
			log.Fatalf("Failed to send BidirectionalStreaming request: %v", err)
		}
		response, err := bidirectionalstream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive BidirectionalStreaming response: %v", err)
		}
		fmt.Printf("BidirectionalStreaming Request: %s\n", request)
		fmt.Printf("BidirectionalStreaming Response: %s\n", response)
	}
	err = bidirectionalstream.CloseSend()
	if err != nil {
		log.Fatalf("Failed to close BidirectionalStreaming stream: %v", err)
	}
	fmt.Println("Exit.........")
}
