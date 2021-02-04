package main

import (
	"log"
	"net"

	"github.com/alexandragrecu/go-gRPC-chat/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Failed to listen on port 4000: %v ", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 4000: %v ", err)
	}
}