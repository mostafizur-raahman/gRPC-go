package main

import (
	"context"
	"fmt"
	"net"

	proto "github.com/mostafizur-raahman/gRPC-go/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedHelloServiceServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":8000")

	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer() // engine server

	proto.RegisterHelloServiceServer(srv, &server{})

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) ServerReply(c context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("Received request ....", r.Request)
	fmt.Println("Send response...")

	return &proto.HelloResponse{}, nil
}
