// Package main implements a server for ChatBot service.
package server

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/bijeshos/grpc_example/protof"
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
	log.Printf("Received from client: %v", in.GetQuery())
	time.Now()
	return &pb.ChatBotReply{Message: "Acknowledement Id : " + strconv.Itoa(rand.Int())}, nil
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
