package storage

import (
	"context"
	billmodel "users/internal/domain/bills/models"
	"users/internal/domain/users/dto"
	"users/internal/domain/users/models"
)

type UsersRep interface {
	CreateUser(ctx context.Context, user models.Users) (models.Users, error)
	GetUser(ctx context.Context, id string) (models.Users, error)
	GetUsers(ctx context.Context) ([]models.Users, error)
	UpdateUsers(ctx context.Context, id string, updateduser *models.Users) (models.Users, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	EnableAccount(ctx context.Context, id string) (string, error)
	LoginUser(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type Billstxn interface {
	InsertTxn(ctx context.Context, txn billmodel.Transaction) error
}
