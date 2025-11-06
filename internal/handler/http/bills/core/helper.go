package core

import (
	"users/internal/domain/bills/dto"
	"users/internal/domain/bills/models"
)

func Tomodelreq(user dto.BillrequestCheck) models.Bills {
	return models.Bills{
		CustomerNumber: user.CustomerNumber,
		ServiceType:    user.ServiceType,
	}
}
