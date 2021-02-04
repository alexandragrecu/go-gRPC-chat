package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/alexandragrecu/go-gRPC-chat/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Buna! Eu sunt clientul!",
	}

	response, err := c.SayHello(context.Background(),&message)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Raspuns de la Server: %s", response.Body)
}

