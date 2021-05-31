package main

import (
	"log"
	"os"

	"github.com/bijeshos/grpc-go-example/client"
	"github.com/bijeshos/grpc-go-example/server"
)

func main() {
	log.Println("starting product inventory server example")
	mode := os.Args[1]

	switch mode {
	case "server":
		server.RunServer()
	case "client":
		client.RunClient()
	default:
		log.Println("Incorrect mode: " + mode)
	}
}
