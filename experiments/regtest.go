package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I am listening")
}

const (
	minTCPPort         = 8081 //0
	maxTCPPort         = 8181 //65535
	maxReservedTCPPort = 1024
	maxRandTCPPort     = maxTCPPort - (maxReservedTCPPort + 1)
)

var (
	tcpPortRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {

	config := api.DefaultConfig()
	config.Address = "localhost:8500"

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	s := strconv.Itoa(randomTCPPort())
	randPort := ":" + s

	// Get a handle to the KV API
	kv := client.KV()
	// PUT a new KV pair
	p := &api.KVPair{Key: "port", Value: []byte(randPort)}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	pair, _, err := kv.Get("port", nil)
	if err != nil {
		fmt.Errorf("Error trying accessing value at key: %v", pair.Key)
	}
	fmt.Println("value: ", string(pair.Key), string(pair.Value))

	err = http.ListenAndServe(randPort, nil)
	if err != nil {
		fmt.Println(err)
	}
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
