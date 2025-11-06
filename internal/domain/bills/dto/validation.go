package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v *BillrequestCheck) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.CustomerNumber, validation.Required, validation.Length(1, 100).Error("customer number is required")),
		validation.Field(&v.ServiceType, validation.Required, validation.In(Electricity, Water, Gas).Error("service type is required")),
	)
}
