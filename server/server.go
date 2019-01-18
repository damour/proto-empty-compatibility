package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s server) Bar(context.Context, *empty.Empty) (*BarResponse, error) {
	fmt.Println("Process request")

	return &BarResponse{Id: "123"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:15444")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := server{}

	grpcServer := grpc.NewServer()
	RegisterFooServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
