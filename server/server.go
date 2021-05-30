// Package server implements a server for ChatBot service.
package server

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/bijeshos/grpc_go_example/protof"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement protof.ChatBotServer.
type server struct {
	pb.UnimplementedChatBotServer
}

// Chat implements protof.ChatBotServer
func (s *server) Chat(ctx context.Context, in *pb.ChatBotRequest) (*pb.ChatBotReply, error) {
	log.Printf("Query received from client: %v", in.GetQuery())

	//This is a rudimentary way of checking input. Not recommanded for real life scenarios.
	//Doing it this way to keeping it simple.
	if strings.Contains(in.GetQuery(), "What is the time now?") {
		return &pb.ChatBotReply{Message: "Current time is " + time.Now().String()}, nil
	} else {
		return &pb.ChatBotReply{Message: "Processing query [ " + in.GetQuery() + " ] is not implemented yet"}, nil
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
	pb.RegisterChatBotServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
