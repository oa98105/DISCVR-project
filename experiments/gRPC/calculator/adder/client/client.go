package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/DISCVR-project/experiments/gRPC/calculator/adder"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdditionClient(conn)

	// Contact the server and print out its response.
	operands := &pb.Operands{A: 1, B: 3}
	r, err := c.Add(context.Background(), operands)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Println(r)
}
