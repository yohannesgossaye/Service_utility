package service

import (
	"context"
	"users/internal/domain/users/dto"
	"users/internal/domain/users/models"
)

type UserS interface {
	CreateUser(ctx context.Context, user models.Users) (models.Users, error)
	GetUser(ctx context.Context, id string) (models.Users, error)
	GetUsers(ctx context.Context) ([]models.Users, error)
	UpdateUsers(ctx context.Context, id string, updateduser *models.Users) (models.Users, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	EnableAccount(ctx context.Context, id string) (string, error)
	LoginUser(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}
