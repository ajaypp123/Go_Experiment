# Introduction
gRPC is an open-source remote procedure call (RPC) framework developed by Google. It allows two or more machines to communicate with each other over a network, using a protocol buffer data format.

gRPC uses Protocol Buffers, a language-agnostic binary serialization format, to define the structure of the data being transmitted between the client and the server. This allows for efficient serialization and deserialization of data, which is important for performance and scalability.

gRPC supports many programming languages including C++, Java, Python, Ruby, Go, and more. It provides many features such as bi-directional streaming, flow control, and blocking or non-blocking calls, and supports both unary and streaming APIs.

One of the key advantages of gRPC is its support for server-side streaming, client-side streaming, and bidirectional streaming. This makes it well-suited for applications that require real-time data streaming, such as chat applications and gaming platforms.

Overall, gRPC provides a modern and efficient way to implement RPC communication between services, with cross-platform and cross-language support, making it an excellent choice for microservices architectures.

# gRPC provides four types of communication patterns:

- Unary RPCs: Unary RPCs are one-to-one request-response messages. The client sends a single request message to the server and the server responds with a single response message.

- Server streaming RPCs: In server streaming RPCs, the client sends a request message to the server and the server responds with a stream of messages. The client reads the stream until there are no more messages.

- Client streaming RPCs: In client streaming RPCs, the client sends a stream of messages to the server and the server responds with a single response message.

- Bidirectional streaming RPCs: In bidirectional streaming RPCs, both the client and the server send a stream of messages to each other. The client can read messages from the server while writing messages to the server at the same time.

# gRPC and Rest Difference

gRPC can be seen as an alternative to REST for building APIs. While REST is based on HTTP/1.1, which is text-based and requires parsing and serialization of data, gRPC is based on HTTP/2, which is binary-based and allows for more efficient data transfer. gRPC uses Protocol Buffers for serialization, which results in smaller message sizes and faster processing compared to JSON used in REST.

Whether or not gRPC can completely replace REST depends on the use case. gRPC is well-suited for use cases where high performance and low latency are critical, such as microservices and IoT devices. On the other hand, REST is still widely used and well-established, especially for web-based APIs.

# WebSockets and gRPC are both used for client-server communication but they differ in several ways:

- Protocol: WebSockets use the WebSocket protocol, which is a full-duplex communication protocol over a single TCP connection. On the other hand, gRPC uses the HTTP/2 protocol.

- Payload: WebSockets are primarily used for sending and receiving text and binary data between the client and server. In contrast, gRPC uses Protocol Buffers, which are a language-agnostic, efficient binary serialization format, to exchange structured data between the client and server.

- API: gRPC is designed to be used with remote procedure calls (RPCs) and provides a mechanism to define services, methods, and messages that are used to exchange data between the client and server. WebSockets, on the other hand, do not have a specific API for defining services and methods.

- Bi-directional: WebSockets allow for bi-directional communication between the client and server, while gRPC also supports bi-directional streaming and server-side streaming.

- Performance: gRPC is designed for high-performance, low-latency communication, making it a good choice for applications that require fast data transfer. WebSockets are also optimized for real-time data transfer, but their performance can be impacted by factors such as latency and network congestion.

# gRPC provides a number of built-in mechanisms for scaling, such as:

1. Load balancing: gRPC provides client-side load balancing, which allows a client to distribute its requests across multiple server instances. This can be done using round-robin, random, or other algorithms.

2. Connection pooling: gRPC allows clients to maintain a pool of connections to servers, which can improve performance by avoiding the overhead of establishing a new connection for each request.
- grpc server listen on port which can have multiple conncetion based on client ip as IP:<PORT>
- On server conection pool can be managed
```
s := grpc.NewServer(
    grpc.KeepaliveParams(keepalive.ServerParameters{
        MaxConnectionIdle:     5 * time.Minute,
        MaxConnectionAge:      30 * time.Minute,
        MaxConnectionAgeGrace: 5 * time.Minute,
        Time:                  1 * time.Minute,
        Timeout:               20 * time.Second,
    }),
)
```

3. Streaming: gRPC supports bidirectional streaming, which allows clients and servers to send and receive a stream of messages. This can be used to reduce the number of requests and responses required, and to enable more efficient use of resources.

4. Async: gRPC supports asynchronous communication, which allows clients to send requests without waiting for a response. This can be used to improve performance and reduce latency, especially in high-traffic scenarios.
- Async client example provided in Project
```
AsyncGetStudent
AsyncGetAllStudent
```
- Async server example provided
```
    // start the server asynchronously
    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

	// now we can also add grpc-proxy for converting to rest api
```

5. In addition to these built-in mechanisms, gRPC can be integrated with other scaling technologies, such as Kubernetes, to provide a more complete solution for managing large-scale distributed systems.

# handle multiple service
```
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &communication.EmployeeServer{})
	pb.RegisterStudentServiceServer(s, &communication.StudentServer{})
    if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
```

# rest combine with grpc
gRPC-Gateway is a tool that serves as a proxy between a gRPC server and a RESTful web client. It generates a reverse-proxy server which translates incoming RESTful HTTP/1.1 requests to gRPC requests and sends them to a gRPC server, and then translates the gRPC responses back to the HTTP/1.1 responses.

gRPC-Gateway allows you to expose a gRPC service as a RESTful API with minimal boilerplate code. It supports all HTTP methods, path parameters, query parameters, request and response bodies, and other RESTful features. It also generates OpenAPI (Swagger) specifications for the exposed RESTful API, which makes it easier to integrate with other systems.

To use gRPC-Gateway, you need to define a gRPC service in a protobuf file and annotate it with special options. You also need to define a separate gateway service in the same protobuf file, which maps the RESTful API to the gRPC API. Once you have defined the service, you can use the protoc-gen-grpc-gateway plugin to generate the gRPC-Gateway code.

Overall, gRPC-Gateway simplifies the process of exposing gRPC services as RESTful APIs, which allows gRPC services to be integrated with a wider range of clients and systems.

```
ClientRest -> (grpc-gateway) -> grpcServer
github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.1
```

```
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
```

# node failure
- Retry if client not recive any response as ultimetly in is req and res.

# Comple proto file

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/runtime

export PATH=$PATH:$GOPATH/bin
source ~/.bashrc

cd proto
# Client Code: --go_out=.
# Server Code: --go-grpc_out=.
# GRPC Proxy: --grpc-gateway_out=.
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. *.proto
```

# Run Application
```
go mod vendor

cd server/
go run main.go

cd client/
go run main.go

# Kill server
netstat -ano | findstr :50051
taskkill /F /PID 14460
```
