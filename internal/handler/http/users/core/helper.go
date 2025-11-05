package core

import (
	"users/internal/domain/dto"
	"users/internal/domain/models"
)

func Tomodelreq(user dto.UsercreateRequest) models.Users {
	return models.Users{
		FullName:     user.FullName,
		Email:        user.Email,
		Phone_number: user.Phone_number,
		Balance:      user.Balance,
	}
}
