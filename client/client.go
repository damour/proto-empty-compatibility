package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:15444", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := NewFooClient(conn)

	req := &empty.Empty{}

	res, err := client.Bar(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Response (with unrecognized fields): ", res)
}
