package service

import (
	"context"
	"log"
	"time"

	pb "users/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCBillClient struct {
	client pb.PaymentServiceClient
}

func NewGRPCBillClient() *GRPCBillClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	client := pb.NewPaymentServiceClient(conn)
	return &GRPCBillClient{client: client}
}

func (g *GRPCBillClient) PayBill(ctx context.Context, req *pb.PayBillRequest) (*pb.PayBillResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return g.client.PayBill(ctx, req)
}
