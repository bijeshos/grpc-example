package main

import (
	"log"
	"os"

	"github.com/bijeshos/grpc_example/client"
	"github.com/bijeshos/grpc_example/server"
)

func main() {
	log.Println("starting chatbot server example")
	mode := os.Args[1]

	if mode == "server" {
		server.RunServer()
	}

	if mode == "client" {
		client.RunClient()
	}

}
