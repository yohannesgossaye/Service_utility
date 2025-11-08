package bills

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"users/internal/domain/bills/dto"
	"users/internal/domain/bills/models"
	middleware "users/internal/handler/middleware"
	"users/internal/storage"
	"users/pkgs/logger"
	pb "users/proto/gen"

	"os"

	gr "users/internal/service/bills/grpc_client"
)

type BillS struct {
	txnRepo    storage.Billstxn
	grpcClient *gr.GRPCBillClient
	log        logger.Logger
}

func NewBills(txnRepo storage.Billstxn, grpcClient *gr.GRPCBillClient, log logger.Logger) *BillS {
	return &BillS{
		txnRepo:    txnRepo,
		grpcClient: grpcClient,
		log:        log,
	}
}

func (s *BillS) GetBills(ctx context.Context, req dto.BillrequestCheck) (models.Bills, error) {
	var bills []models.Bills
	var result models.Bills

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

	userinfo, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		return dto.BillPaymentResponse{}, fmt.Errorf("failed to get user from context: %v", err)
	}

	// Call gRPC to update MockAPI.io bill status
	grpcReq := &pb.PayBillRequest{
		CustomerNumber: req.CustomerNumber,
		ServiceType:    req.ServiceType,
		AmountDue:      float64(req.Amount),
	}

	res, err := b.grpcClient.PayBill(ctx, grpcReq)
	if err != nil {
		return dto.BillPaymentResponse{}, fmt.Errorf("failed to process payment via gRPC: %v", err)
	}

	// Save transaction to MongoDB
	transaction := models.Transaction{
		TransactionID:  res.TransactionId,
		UserID:         userinfo.UserID,
		CustomerNumber: req.CustomerNumber,
		ServiceType:    req.ServiceType,
		Amount:         req.Amount,
		Status:         res.Status,
		Timestamp:      time.Now(),
	}

	if err := b.txnRepo.InsertTxn(ctx, transaction); err != nil {
		b.log.Infof("Failed to save transaction to database: %v", err)
	}

	b.log.Infof("Transaction saved: TransactionID=%s, UserID=%s, Amount=%.2f",
		res.TransactionId, userinfo.UserID, req.Amount)

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
