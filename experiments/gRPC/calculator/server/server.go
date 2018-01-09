package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/DISCVR-project/experiments/gRPC/calculator"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.Operands) (*pb.Response, error) {
	a, b := in.A, in.B
	return &pb.Response{Result: (a + b)}, nil
}

func (s *server) Div(ctx context.Context, in *pb.Operands) (*pb.Response, error) {
	a, b := in.A, in.B
	return &pb.Response{Result: (a / b)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	s.Serve(lis)
}
