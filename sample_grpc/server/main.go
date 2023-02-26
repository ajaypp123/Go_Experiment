package main

import (
	"log"
	"net"
	"time"

	pb "github.com/ajaypp123/sample_grpc/proto/service"
	"github.com/ajaypp123/sample_grpc/server/communication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	grpcPort  = ":50051"     // the port for the gRPC server
	gwPort    = ":8081"      // the port for the gRPC reverse proxy server
	endpoints = "/employee/" // the API endpoints for the gRPC reverse proxy server
)

func main() {

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     5 * time.Minute,
			MaxConnectionAge:      30 * time.Minute,
			MaxConnectionAgeGrace: 5 * time.Minute,
			Time:                  1 * time.Minute,
			Timeout:               20 * time.Second,
		}),
	)

	empServer := &communication.EmployeeServer{}
	empServer.Init()
	pb.RegisterEmployeeServiceServer(s, empServer)

	stdServer := &communication.StudentServer{}
	stdServer.Init()
	pb.RegisterStudentServiceServer(s, stdServer)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	/*
		// create a gRPC server
		lis, err := net.Listen("tcp", grpcPort)
		grpcServer := grpc.NewServer()
		empServer := &communication.EmployeeServer{}
		empServer.Init()
		pb.RegisterEmployeeServiceServer(grpcServer, empServer)

		stdServer := &communication.StudentServer{}
		stdServer.Init()
		pb.RegisterStudentServiceServer(grpcServer, stdServer)

		// start the server asynchronously
		go func() {
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()

		// create a Gorilla Mux router
		gwMux := runtime.NewServeMux()

		// register the gRPC reverse proxy handlers
		err = pb.RegisterEmployeeServiceHandlerFromEndpoint(
			context.Background(),
			gwMux,
			grpcPort,
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			log.Fatalf("failed to register gateway: %v", err)
		}

		// create a Gorilla Mux router and serve the reverse proxy
		muxRouter := mux.NewRouter()
		muxRouter.Handle(endpoints, gwMux)

		log.Printf("gRPC server listening on %s", grpcPort)
		log.Printf("gRPC gateway listening on %s", gwPort)
		log.Fatal(http.ListenAndServe(gwPort, muxRouter))

		// wait for the server to finish (e.g. on interrupt signal)
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch

		// stop the server gracefully
		s.GracefulStop()
	*/
}
