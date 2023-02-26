package communication

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ajaypp123/sample_grpc/proto/service"
)

type EmployeeServer struct {
	pb.UnimplementedEmployeeServiceServer
}

func (s *EmployeeServer) Init() {
	initEmpDB()
	fmt.Println("=================================")
	fmt.Println()
}

// ######################### Unary ##########################
// Send single employee
func (s *EmployeeServer) GetEmployee(ctx context.Context, in *pb.GetEmployeeRequest) (*pb.EmployeeResponse, error) {
	fmt.Println()
	fmt.Println("============== Unary ===================")
	fmt.Println("Entry.........")
	resp := &pb.EmployeeResponse{
		Employee: empMap[in.GetEmployeeId()],
	}
	fmt.Printf("GetEmployee Unary Request: %s\n", in)
	fmt.Printf("GetEmployee Unary Response: %s\n", resp)
	fmt.Println("Exit.........")
	return resp, nil
}

// ######################### Server Streaming ##########################
// Send employee stream
func (s *EmployeeServer) GetAllEmployees(in *pb.GetAllEmployeeRequest, stream pb.EmployeeService_GetAllEmployeesServer) error {
	fmt.Println()
	fmt.Println("============== Server Streaming ===================")
	fmt.Println("Entry.........")
	fmt.Printf("GetAllEmployees Server Streaming Request: %s\n", in)
	for _, id := range in.GetEmployeeIds() {
		res := &pb.EmployeeResponse{
			Employee: empMap[id],
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Failed to call GetAllEmployees: %v", err)
			return err
		}
		fmt.Printf("GetAllEmployees Server Streaming Response: %s\n", res)
	}
	fmt.Println("Exit.........")
	return nil
}

// ######################### Client Streaming ##########################
// Add employee stream
func (s *EmployeeServer) AddEmployees(stream pb.EmployeeService_AddEmployeesServer) error {
	fmt.Println()
	fmt.Println("============== Client Streaming ===================")
	fmt.Println("Entry.........")
	var emps []*pb.Employee
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to call AddEmployees: %v", err)
			return err
		}
		fmt.Printf("AddEmployees ClientStreaming Request: %v\n", req)
		employee := req.GetEmployee()
		empMap[employee.EmployeeId] = employee
		emps = append(emps, employee)
		/*
			if len(emps) == 3 {
				break
			}
		*/
	}

	res := &pb.EmployeesResponse{
		Employees: emps,
	}
	stream.SendAndClose(res)
	fmt.Printf("AddEmployees ClientStreaming Response: %v\n", res)
	fmt.Println("Exit.........")
	return nil
}

// ######################### Bidirectional Streaming ##########################
func (s *EmployeeServer) Chat(stream pb.EmployeeService_ChatServer) error {
	fmt.Println()
	fmt.Println("============== Bidirectional Streaming ===================")
	fmt.Println("Entry.........")
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to call BidirectionalStreaming: %v", err)
			return err
		}
		fmt.Printf("Chat BidirectionalStreaming Request: %v\n", req)
		employee, ok := empMap[req.GetEmployeeId()]
		if !ok {
			//return io.EOF
			return nil
		}
		res := &pb.EmployeeResponse{
			Employee: employee,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Failed to call BidirectionalStreaming: %v", err)
			return err
		}
		fmt.Printf("Chat BidirectionalStreaming Response: %v\n", res)
		fmt.Println("Exit.........")
	}
}
