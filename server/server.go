package main

import (
	"fmt"
	"log"
	"net"

	"github.com/NikhilSharma03/expensetracker/server/expensepb"
	"github.com/NikhilSharma03/expensetracker/server/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Something went wrong", err.Error())
	}

	gRPCServer := grpc.NewServer()
	expensepb.RegisterExpenseServiceServer(gRPCServer, &service.ExpenseServer{})
	fmt.Println("Starting gRPC Server...")
	if gRPCServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err.Error())
	}
}
