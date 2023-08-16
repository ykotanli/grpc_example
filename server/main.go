package main

import (
	"context"
	"net"

	"github.com/ykotanli/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedAddServiceServer // Embed the UnimplementedAddServiceServer struct to implement the interface)
}

// Implement the methods from the AddServiceServer interface
func (*server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil

}

func (*server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{}) // Register the server
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
