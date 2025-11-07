package main

import (
	"fmt"
	"log"
	"net"

	"users/internal/handler/grpc/server"
	payment "users/proto/gen"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServiceServer(grpcServer, &server.PaymentServer{})

	fmt.Println("🚀 gRPC server running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
