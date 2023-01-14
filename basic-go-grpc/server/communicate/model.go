package communicate

import pb "github.com/ajaypp/basic-go-grpc/proto"

// this is the struct to be created, pb is imported upstairs
type HelloServer struct {
	pb.GreetServiceServer
}
