// Package server implements a server for Product service.
package server

import (
	"context"
	"log"
	"net"
	"strings"

	pb "github.com/bijeshos/grpc-go-example/protof"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement protof.ProductServer.
type server struct {
	pb.UnimplementedProductServer
}

// Search implements protof.ProductServer
func (s *server) Search(ctx context.Context, in *pb.ProductSearchRequest) (*pb.ProductSearchReply, error) {
	log.Printf("Product received from client: %v", in.GetProduct())

	//This is a rudimentary way of checking input. Not recommanded for real life scenarios.
	//Doing it this way to keeping it simple.
	if strings.Contains(in.GetProduct(), "Laptop") {
		return &pb.ProductSearchReply{Message: in.GetProduct() + " is in stock"}, nil
	} else {
		return &pb.ProductSearchReply{Message: in.GetProduct() + " is not in stock"}, nil
	}

}

func RunServer() {
	log.Println("initiating server")
	lis, err := net.Listen("tcp", port)
	log.Println("listening to tcp port")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("started new grpc server")
	pb.RegisterProductServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
