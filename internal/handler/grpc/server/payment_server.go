package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	payment "users/proto/gen"
)

type PaymentServer struct {
	payment.UnimplementedPaymentServiceServer
}

func (s *PaymentServer) PayBill(ctx context.Context, req *payment.PayBillRequest) (*payment.PayBillResponse, error) {
	apiURL := GetBillAPI(req.ServiceType)
	if apiURL == "" {
		return nil, fmt.Errorf("API URL not configured for service type: %s", req.ServiceType)
	}

	billID, err := findBillByCustomerNumber(apiURL, req.CustomerNumber, req.ServiceType)
	if err != nil {
		return nil, fmt.Errorf("failed to find bill: %v", err)
	}

	payload := map[string]interface{}{
		"status": "paid",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	client := &http.Client{}
	updateURL := fmt.Sprintf("%s/%s", apiURL, billID)

	httpReq, err := http.NewRequest(http.MethodPut, updateURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("external API returned status %d for URL: %s", resp.StatusCode, updateURL)
	}

	// Return success response
	return &payment.PayBillResponse{
		TransactionId: fmt.Sprintf("TXN-%s", req.CustomerNumber),
		Message:       "Bill paid successfully",
		Status:        "success",
	}, nil
}

// findBillByCustomerNumber searches for a bill by customer_number and returns its ID
func findBillByCustomerNumber(apiURL, customerNumber, serviceType string) (string, error) {
	client := &http.Client{}
	resp, err := client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch bills: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("external API returned status %d", resp.StatusCode)
	}

	var bills []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&bills); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if len(bills) == 0 {
		return "", fmt.Errorf("no bills returned from API: %s", apiURL)
	}

	// Search for the bill with matching customer_number and service_type
	for _, bill := range bills {
		// Get customer_number
		custNumVal, custNumExists := bill["customer_number"]
		if !custNumExists {
			continue
		}
		custNum, custNumOk := custNumVal.(string)
		if !custNumOk {
			continue
		}

		// Get service_type
		svcTypeVal, svcTypeExists := bill["service_type"]
		if !svcTypeExists {
			continue
		}
		svcType, svcTypeOk := svcTypeVal.(string)
		if !svcTypeOk {
			continue
		}

		// Check if both customer_number and service_type match
		if custNum == customerNumber && svcType == serviceType {
			// Get the ID field (MockAPI.io uses "id" as the primary key)
			idVal, idExists := bill["id"]
			if !idExists {
				return "", fmt.Errorf("found matching bill but ID field is missing for customer_number: %s", customerNumber)
			}

			// Handle different ID types (MockAPI.io can return ID as string or number)
			switch id := idVal.(type) {
			case string:
				return id, nil
			case float64:
				// JSON numbers decode as float64
				return fmt.Sprintf("%.0f", id), nil
			case int:
				return fmt.Sprintf("%d", id), nil
			case int64:
				return fmt.Sprintf("%d", id), nil
			default:
				return "", fmt.Errorf("found matching bill but ID has unexpected type: %T for customer_number: %s", id, customerNumber)
			}
		}
	}

	return "", fmt.Errorf("bill not found for customer_number: %s with service_type: %s", customerNumber, serviceType)
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
