package storage

import (
	"context"
	"users/internal/domain/models"
)

type UsersRep interface {
	CreateUser(ctx context.Context, user models.Users) (models.Users, error)
	GetUser(ctx context.Context, id string) (models.Users, error)
	GetUsers(ctx context.Context) ([]models.Users, error)
	UpdateUsers(ctx context.Context, id string, updateduser *models.Users) (models.Users, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	EnableAccount(ctx context.Context, id string) (string, error)
}
