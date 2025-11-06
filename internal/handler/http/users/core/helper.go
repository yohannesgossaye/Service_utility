package core

import (
	"users/internal/domain/users/dto"
	"users/internal/domain/users/models"
)

func Tomodelreq(user dto.UsercreateRequest) models.Users {
	return models.Users{
		FullName:     user.FullName,
		Email:        user.Email,
		Phone_number: user.Phone_number,
		Balance:      user.Balance,
		Password:     user.Password,
	}
}
