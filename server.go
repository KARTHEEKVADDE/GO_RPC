package main

import (
	"log"
	"net"

	"github.com/kartheekvadde/GO_RPC/chat"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("Failed to listen on port 12345: %v", err)
	}

	s := chat.Server{}
	
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 12345: %v", err)
	}

}