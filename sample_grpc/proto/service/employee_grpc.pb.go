// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: employee.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	// Unary RPC: Send a single request and receive a single response.
	GetEmployee(ctx context.Context, in *GetEmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	// Server Streaming RPC: Send a single request and receive multiple responses.
	GetAllEmployees(ctx context.Context, in *GetAllEmployeeRequest, opts ...grpc.CallOption) (EmployeeService_GetAllEmployeesClient, error)
	// Client Streaming RPC: Send multiple requests and receive a single response.
	AddEmployees(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddEmployeesClient, error)
	// Bidirectional Streaming RPC: Send and receive multiple requests and responses.
	Chat(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_ChatClient, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetEmployee(ctx context.Context, in *GetEmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/service.EmployeeService/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetAllEmployees(ctx context.Context, in *GetAllEmployeeRequest, opts ...grpc.CallOption) (EmployeeService_GetAllEmployeesClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmployeeService_ServiceDesc.Streams[0], "/service.EmployeeService/GetAllEmployees", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetAllEmployeesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetAllEmployeesClient interface {
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceGetAllEmployeesClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetAllEmployeesClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) AddEmployees(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddEmployeesClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmployeeService_ServiceDesc.Streams[1], "/service.EmployeeService/AddEmployees", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceAddEmployeesClient{stream}
	return x, nil
}

type EmployeeService_AddEmployeesClient interface {
	Send(*AddEmployeeRequest) error
	CloseAndRecv() (*EmployeesResponse, error)
	grpc.ClientStream
}

type employeeServiceAddEmployeesClient struct {
	grpc.ClientStream
}

func (x *employeeServiceAddEmployeesClient) Send(m *AddEmployeeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceAddEmployeesClient) CloseAndRecv() (*EmployeesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmployeesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmployeeService_ServiceDesc.Streams[2], "/service.EmployeeService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceChatClient{stream}
	return x, nil
}

type EmployeeService_ChatClient interface {
	Send(*GetEmployeeRequest) error
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceChatClient struct {
	grpc.ClientStream
}

func (x *employeeServiceChatClient) Send(m *GetEmployeeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceChatClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	// Unary RPC: Send a single request and receive a single response.
	GetEmployee(context.Context, *GetEmployeeRequest) (*EmployeeResponse, error)
	// Server Streaming RPC: Send a single request and receive multiple responses.
	GetAllEmployees(*GetAllEmployeeRequest, EmployeeService_GetAllEmployeesServer) error
	// Client Streaming RPC: Send multiple requests and receive a single response.
	AddEmployees(EmployeeService_AddEmployeesServer) error
	// Bidirectional Streaming RPC: Send and receive multiple requests and responses.
	Chat(EmployeeService_ChatServer) error
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) GetEmployee(context.Context, *GetEmployeeRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) GetAllEmployees(*GetAllEmployeeRequest, EmployeeService_GetAllEmployeesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllEmployees not implemented")
}
func (UnimplementedEmployeeServiceServer) AddEmployees(EmployeeService_AddEmployeesServer) error {
	return status.Errorf(codes.Unimplemented, "method AddEmployees not implemented")
}
func (UnimplementedEmployeeServiceServer) Chat(EmployeeService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.EmployeeService/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, req.(*GetEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetAllEmployees_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllEmployeeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetAllEmployees(m, &employeeServiceGetAllEmployeesServer{stream})
}

type EmployeeService_GetAllEmployeesServer interface {
	Send(*EmployeeResponse) error
	grpc.ServerStream
}

type employeeServiceGetAllEmployeesServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetAllEmployeesServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_AddEmployees_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).AddEmployees(&employeeServiceAddEmployeesServer{stream})
}

type EmployeeService_AddEmployeesServer interface {
	SendAndClose(*EmployeesResponse) error
	Recv() (*AddEmployeeRequest, error)
	grpc.ServerStream
}

type employeeServiceAddEmployeesServer struct {
	grpc.ServerStream
}

func (x *employeeServiceAddEmployeesServer) SendAndClose(m *EmployeesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceAddEmployeesServer) Recv() (*AddEmployeeRequest, error) {
	m := new(AddEmployeeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmployeeService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).Chat(&employeeServiceChatServer{stream})
}

type EmployeeService_ChatServer interface {
	Send(*EmployeeResponse) error
	Recv() (*GetEmployeeRequest, error)
	grpc.ServerStream
}

type employeeServiceChatServer struct {
	grpc.ServerStream
}

func (x *employeeServiceChatServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceChatServer) Recv() (*GetEmployeeRequest, error) {
	m := new(GetEmployeeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployee",
			Handler:    _EmployeeService_GetEmployee_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllEmployees",
			Handler:       _EmployeeService_GetAllEmployees_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddEmployees",
			Handler:       _EmployeeService_AddEmployees_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _EmployeeService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "employee.proto",
}
