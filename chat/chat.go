package chat

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	log.Printf("Received message body %s from client: %s", message.Body, message.Name)

	return &Message{Body: fmt.Sprintf("Hello %s from the Server!", message.Name)}, nil
}

func (s *Server) Addition(ctx context.Context, req *Request) (*Response, error) {

	log.Printf("Received message body for Addition:")

	result := req.GetFirst() + req.GetSecond()
	return &Response{Result: result}, nil
}

func (s *Server) Multiplication(ctx context.Context, req *Request) (*Response, error) {

	log.Printf("Received message body for Multiplication:")

	result := req.GetFirst() * req.GetSecond()
	return &Response{Result: result}, nil
}

func (s *Server) Subtraction(ctx context.Context, req *Request) (*Response, error) {

	log.Printf("Received message body for Subtraction:")

	result := req.GetFirst() - req.GetSecond()
	return &Response{Result: result}, nil
}

func (s *Server) Division(ctx context.Context, req *Request) (*Response, error) {

	log.Printf("Received message body for Division:")

	result := req.GetFirst() / req.GetSecond()
	return &Response{Result: result}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
}
