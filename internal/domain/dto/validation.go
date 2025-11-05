package dto

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phone_numberRegex = regexp.MustCompile(`^(09|\+2519)\d{8}$`)
)

func (v UsercreateRequest) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.FullName, validation.Required.Error("fullname is required")),
		validation.Field(&v.Email, validation.Required.Error("email is required"),
			validation.Match(emailRegex).Error("invalid email provided "),
		),
		validation.Field(&v.Phone_number, validation.Required.Error("phone number is required"),
			validation.Match(phone_numberRegex).Error("invalid phone number provided "),
		),
		validation.Field(&v.Balance, validation.Required.Error("balance is required"),
			validation.Min(float64(100)).Error("minimum balance is 100 birr"),
		),
	)
}
