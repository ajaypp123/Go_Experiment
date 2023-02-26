package communication

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ajaypp123/sample_grpc/proto/service"
)

type StudentServer struct {
	pb.UnimplementedStudentServiceServer
}

func (s *StudentServer) Init() {
	initStdDB()
	fmt.Println("=================================")
	fmt.Println()
}

// ######################### Unary ##########################
// Send single student
func (s *StudentServer) GetStudent(ctx context.Context, in *pb.GetStudentRequest) (*pb.StudentResponse, error) {
	fmt.Println()
	fmt.Println("============== Unary ===================")
	fmt.Println("Entry.........")
	resp := &pb.StudentResponse{
		Student: stdMap[in.GetStudentId()],
	}
	fmt.Printf("GetStudent Unary Request: %s\n", in)
	fmt.Printf("GetStudent Unary Response: %s\n", resp)
	fmt.Println("Exit.........")
	return resp, nil
}

// ######################### Server Streaming ##########################
// Send student stream
func (s *StudentServer) GetAllStudents(in *pb.GetAllStudentRequest, stream pb.StudentService_GetAllStudentsServer) error {
	fmt.Println()
	fmt.Println("============== Server Streaming ===================")
	fmt.Println("Entry.........")
	fmt.Printf("GetAllStudents Server Streaming Request: %s\n", in)
	for _, id := range in.GetStudentIds() {
		res := &pb.StudentResponse{
			Student: stdMap[id],
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Failed to call GetAllStudents: %v", err)
			return err
		}
		fmt.Printf("GetAllStudents Server Streaming Response: %s\n", res)
	}
	fmt.Println("Exit.........")
	return nil
}

// ######################### Client Streaming ##########################
// Add student stream
func (s *StudentServer) AddStudents(stream pb.StudentService_AddStudentsServer) error {
	fmt.Println()
	fmt.Println("============== Client Streaming ===================")
	fmt.Println("Entry.........")
	var emps []*pb.Student
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to call AddStudents: %v", err)
			return err
		}
		fmt.Printf("AddStudents ClientStreaming Request: %v\n", req)
		student := req.GetStudent()
		stdMap[student.StudentId] = student
		emps = append(emps, student)
		/*
			if len(emps) == 3 {
				break
			}
		*/
	}

	res := &pb.StudentsResponse{
		Students: emps,
	}
	stream.SendAndClose(res)
	fmt.Printf("AddStudents ClientStreaming Response: %v\n", res)
	fmt.Println("Exit.........")
	return nil
}

// ######################### Bidirectional Streaming ##########################
func (s *StudentServer) Chat(stream pb.StudentService_ChatServer) error {
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
		student, ok := stdMap[req.GetStudentId()]
		if !ok {
			//return io.EOF
			return nil
		}
		res := &pb.StudentResponse{
			Student: student,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Failed to call BidirectionalStreaming: %v", err)
			return err
		}
		fmt.Printf("Chat BidirectionalStreaming Response: %v\n", res)
		fmt.Println("Exit.........")
	}
}
