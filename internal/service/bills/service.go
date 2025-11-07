package bills

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"users/internal/domain/bills/dto"
	"users/internal/domain/bills/models"
	"users/pkgs/logger"
	pb "users/proto/gen"

	"os"

	gr "users/internal/service/bills/grpc_client"
)

type BillS struct {
	grpcClient *gr.GRPCBillClient
	log        logger.Logger
}

func NewBills(grpcClient *gr.GRPCBillClient, log logger.Logger) *BillS {
	return &BillS{
		grpcClient: grpcClient,
		log:        log,
	}
}

func (s *BillS) GetBills(ctx context.Context, req dto.BillrequestCheck) (models.Bills, error) {
	var bills []models.Bills
	var result models.Bills

	// fetch data from external mock API
	resp, err := http.Get(GetBillAPI(req.ServiceType))
	if err != nil {
		return result, fmt.Errorf("failed to fetch bills: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("external API returned status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&bills); err != nil {
		return result, fmt.Errorf("failed to decode response: %v", err)
	}

	// find the matching bill
	for _, b := range bills {
		if b.CustomerNumber == req.CustomerNumber {
			result = b
			break
		}
	}

	if result.CustomerNumber == "" {
		return result, errors.New("no bill found for this customer number")
	}

	return result, nil
}

func (b *BillS) PayBills(ctx context.Context, req dto.BillPaymentRequest) (dto.BillPaymentResponse, error) {
	grpcReq := &pb.PayBillRequest{
		CustomerNumber: req.CustomerNumber,
		ServiceType:    req.ServiceType,
		AmountDue:      float64(req.Amount),
	}

	res, err := b.grpcClient.PayBill(ctx, grpcReq)
	if err != nil {
		return dto.BillPaymentResponse{}, err
	}
	// we will write after this in mong transaction collection for payment bills and track who pai by user id by middleware
	// You can also save to the transaction table here after success
	return dto.BillPaymentResponse{
		TransactionId: res.TransactionId,
		Message:       res.Message,
		Status:        res.Status,
	}, nil
}

func GetBillAPI(serviceType string) string {
	switch serviceType {
	case "electric":
		return os.Getenv("ELECTRIC_BILL_API")
	case "water":
		return os.Getenv("WATER_BILL_API")
	default:
		return ""
	}
}
