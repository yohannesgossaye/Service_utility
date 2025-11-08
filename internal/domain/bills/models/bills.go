package models

import "time"

type Bills struct {
	FullName       string  `json:"fullname"`
	CustomerNumber string  `json:"customer_number"`
	ServiceType    string  `json:"service_type"`
	Status         string  `json:"status"`
	Amount_due     float64 `json:"amount_due"`
	Due_date       string  `json:"due_date"`
}

type BillPaymentResponse struct {
	TransactionId string `json:"transaction_id"`
	Message       string `json:"message"`
	Status        string `json:"status"`
}

type Transaction struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitempty"`
	TransactionID  string    `bson:"transaction_id" json:"transaction_id"`
	UserID         string    `bson:"user_id" json:"user_id"`
	CustomerNumber string    `bson:"customer_number" json:"customer_number"`
	ServiceType    string    `bson:"service_type" json:"service_type"`
	Amount         float64   `bson:"amount" json:"amount"`
	Status         string    `bson:"status" json:"status"`
	Timestamp      time.Time `bson:"timestamp" json:"timestamp"`
}
