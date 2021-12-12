package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Something went wrong", err.Error())
	}

	gRPCServer := grpc.NewServer()

	if gRPCServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err.Error())
	}
}
