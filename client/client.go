package main

import (
	"context"

	proto "github.com/mostafizur-raahman/gRPC-go/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewHelloServiceClient(conn)

	req := &proto.HelloRequest{Request: "Hello from client"}

	client.ServerReply(context.TODO(), req)
}
