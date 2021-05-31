// Package client implements a client for Product service.
package client

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/bijeshos/grpc-go-example/protof"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultProduct = "Laptop"
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
	c := pb.NewProductClient(conn)

	// Contact the server and print out its response.
	product := defaultProduct
	if len(os.Args) > 2 {
		product = os.Args[2]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Sending product name: %s", product)
	r, err := c.Search(ctx, &pb.ProductSearchRequest{Product: product})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
