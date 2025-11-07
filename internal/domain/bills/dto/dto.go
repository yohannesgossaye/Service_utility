package dto

type BillrequestCheck struct {
	CustomerNumber string `json:"customer_number"`
	ServiceType    string `json:"service_type"`
}

const (
	Electricity string = "electric"
	Water       string = "water"
	Gas         string = "gas"
)

type BillCheckResponse struct {
	FullName       string  `json:"fullname"`
	CustomerNumber string  `json:"customer_number"`
	ServiceType    string  `json:"service_type"`
	Status         string  `json:"status"`
	Amount_due     float64 `json:"amount_due"`
	Due_date       string  `json:"due_date"`
}

type BillPaymentRequest struct {
	CustomerNumber string  `json:"customer_number"`
	ServiceType    string  `json:"service_type"`
	Amount         float64 `json:"amount"`
}

type BillPaymentResponse struct {
	TransactionId string `json:"transaction_id"`
	Message       string `json:"message"`
	Status        string `json:"status"`
}
