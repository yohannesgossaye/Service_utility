package models

import "time"

type Users struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitempty"`
	FullName       string    `bson:"fullname" json:"fullname"`
	Account_number string    `bson:"account_number" json:"account_number"`
	Password       string    `bson:"password" json:"password"`
	Phone_number   string    `bson:"phone_number" json:"phone_number"`
	Balance        float64   `bson:"balance" json:"balance"`
	Is_active      bool      `bson:"is_active" json:"is_active"`
	Email          string    `bson:"email" json:"email"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
}
