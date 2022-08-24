package main

import (
	// "context"
	"log"
	"net/http"
	"github.com/kartheekvadde/GO_RPC/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/gorilla/mux"

)

func main() {

	var conn *grpc.ClientConn
	//grpc.WithInsecure is deprecated: use WithTransportCredentials and insecure.NewCredentials() instead.
	// Will be supported throughout 1.x.  (SA1019)go-staticcheck

	conn, err := grpc.Dial(":12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	
	callHello(c)
	callAddition(c)
	callSubtraction(c)
	callMultiplication(c)
	callDivision(c)
}

func callHello(c chat.ChatServiceClient) {

	message := chat.Message{
		Body: "Hello from Client!",
		Name: "Kartheek Vadde",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Responce from Server: %s", response.Body)
}

func callAddition(c chat.ChatServiceClient) {

	request := chat.Request{
		First:  1345.34,
		Second: 231450.45,
	}

	response, err := c.Addition(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Responce from Server: %v", response.Result)
}

func callSubtraction(c chat.ChatServiceClient) {

	request := chat.Request{
		First:  1345.34,
		Second: 230.45,
	}

	response, err := c.Subtraction(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Responce from Server: %v", response.Result)
}

func callMultiplication(c chat.ChatServiceClient) {

	request := chat.Request{
		First:  45.34,
		Second: 23.45,
	}

	response, err := c.Multiplication(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Responce from Server: %v", response.Result)
}

func callDivision(c chat.ChatServiceClient) {

	request := chat.Request{
		First:  1345.34,
		Second: 20.45,
	}

	response, err := c.Division(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Responce from Server: %v", response.Result)
}
