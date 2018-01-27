package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/DISCVR-project/experiments/gRPC/calculator/multiplier"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

var (
	tcpPortRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

const (
	minTCPPort         = 50051 //0
	maxTCPPort         = 50099 //65535
	maxReservedTCPPort = 1024
	maxRandTCPPort     = maxTCPPort - (maxReservedTCPPort + 1)
)

type server struct{}

func (s *server) Mul(ctx context.Context, in *pb.Operands) (*pb.Response, error) {
	a, b := in.A, in.B
	return &pb.Response{Result: (a * b)}, nil
}

func main() {
	//---------------------- select a randon port --------------------------//
	p := strconv.Itoa(randomTCPPort())
	randPort := ":" + p
	//---------------------- get a handle for consul store -----------------//
	config := api.DefaultConfig()
	config.Address = "localhost:8500"

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()
	// ----------------------------Register server in consul store ---------//
	kvpair := &api.KVPair{Key: "Mul", Value: []byte(randPort)}
	_, err = kv.Put(kvpair, nil)
	if err != nil {
		panic(err)
	}
	//---------------------------- start listening -------------------------//
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", randPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMultiplicationServer(s, &server{})
	s.Serve(lis)
}

func randomTCPPort() int {
	for i := maxReservedTCPPort; i < maxTCPPort; i++ {
		p := tcpPortRand.Intn(maxRandTCPPort) + maxReservedTCPPort + 1
		if isTCPPortAvailable(p) {
			return p
		}
	}
	return -1
}

func isTCPPortAvailable(port int) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	conn, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
