package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	hello "github.com/rosspatil/gRPC-example/server"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(cxt context.Context, r *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println("hello " + r.GetName())
	return &hello.HelloResponse{}, nil
}
