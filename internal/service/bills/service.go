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

	"os"
)

type BillS struct {
	log logger.Logger
}

func NewBills(log logger.Logger) *BillS {
	return &BillS{
		log: log,
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
