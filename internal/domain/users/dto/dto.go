package dto

import "time"

type UsercreateRequest struct {
	FullName     string  `json:"fullname"`
	Email        string  `json:"email"`
	Phone_number string  `json:"phone_number"`
	Balance      float64 `json:"balance"`
	Password     string  `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type UserResponse struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitempty"`
	FullName       string    `bson:"fullname" json:"fullname"`
	Account_number string    `bson:"account_number" json:"account_number"`
	Email          string    `bson:"email" json:"email"`
	Phone_number   string    `bson:"phone_number" json:"phone_number"`
	Balance        float64   `bson:"balance" json:"balance"`
	Is_active      bool      `bson:"is_active" json:"is_active"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
}
