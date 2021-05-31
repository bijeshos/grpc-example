package main

import (
	"log"
	"os"

	"github.com/bijeshos/grpc_go_example/client"
	"github.com/bijeshos/grpc_go_example/server"
)

func main() {
	log.Println("starting product inventory server example")
	mode := os.Args[1]

	if mode == "server" {
		server.RunServer()
	}

	if mode == "client" {
		client.RunClient()
	}

}
