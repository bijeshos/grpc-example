// Package client implements a client for ChatBot service.
package client

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/bijeshos/grpc_go_example/protof"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultQuery = "What is the time now?"
)

func RunClient() {
	log.Println("initiating client")
	// Set up a connection to the server.
	log.Println("connecting  to server")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("creating a new client")
	c := pb.NewChatBotClient(conn)

	// Contact the server and print out its response.
	query := defaultQuery
	if len(os.Args) > 2 {
		query = os.Args[2]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Sending query: %s", query)
	r, err := c.Chat(ctx, &pb.ChatBotRequest{Query: query})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
