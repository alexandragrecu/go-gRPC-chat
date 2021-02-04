package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Am primit mesajul de la client, acesta este: %s", in.Body)
	return &Message{Body: "Salut! eu sunt serverul!"}, nil
}
