package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	hello "github.com/rosspatil/gRPC-example/server"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial Failed: %v", err)
	}
	client := hello.NewHelloServiceClient(conn)
	rs, err := client.Compute(context.Background(), &hello.HelloRequest{Name: "roshan"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs.GetName())
}
