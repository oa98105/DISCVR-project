package main

import (
	"fmt"
	"log"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
)

func init() {
	// Register consul store to libkv
	consul.Register()
}

func main() {
	whatever()
}

func whatever() {
	client := "localhost:8500"

	// Initialize a new store with consul
	kv, err := libkv.NewStore(
		store.CONSUL, // or "consul"
		[]string{client},
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatal("Cannot create store consul")
	}
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	for k, v := range m {
		err = kv.Put(k, []byte(v), &store.WriteOptions{TTL: 1 * time.Minute})
		if err != nil {
			fmt.Errorf("Error trying to put value at key: %v", k)
		}
		pair, err := kv.Get(k)
		if err != nil {
			fmt.Errorf("Error trying accessing value at key: %v", k)
		}

		fmt.Println("value: ", string(pair.Key), string(pair.Value))
	}

	for k := range m {
		if k == "k2" {
			err = kv.Delete(k)
			if err != nil {
				fmt.Errorf("Error trying to delete key %v", k)
			}
		}

	}

}
