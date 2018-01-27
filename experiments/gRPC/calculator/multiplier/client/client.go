package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/DISCVR-project/experiments/gRPC/calculator/multiplier"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

// const (
// 	address = "localhost:50052"
// )

func main() {

	config := api.DefaultConfig()
	config.Address = "localhost:8500"

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// Get a handle to the KV API
	kv := client.KV()
	pair, _, err := kv.Get(os.Args[1], nil)
	if err != nil {
		fmt.Errorf("Error trying accessing value at key: %v", pair.Key)
	}
	address := "localhost" + string(pair.Value)
	fmt.Println(address)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMultiplicationClient(conn)

	// Contact the server and print out its response.
	operands := &pb.Operands{A: 1, B: 3}
	r, err := c.Mul(context.Background(), operands)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Println(r)
}
