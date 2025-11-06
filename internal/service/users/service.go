package users

import (
	"context"
	"fmt"
	"users/internal/domain/users/dto"
	"users/internal/domain/users/models"
	"users/internal/service"
	"users/pkgs/logger"
	"users/pkgs/utils/email"
)

type ServU struct {
	SS    service.UserS
	log   logger.Logger
	email email.Sender
}

func InitSvc(ss service.UserS, log logger.Logger, email email.Sender) *ServU {
	return &ServU{
		SS:    ss,
		log:   log,
		email: email,
	}
}

func (s *ServU) CreateUser(ctx context.Context, user models.Users) (models.Users, error) {
	create, err := s.SS.CreateUser(ctx, user)
	if err != nil {
		return models.Users{}, nil
	}
	go s.SendEmail(create.Email, create.FullName)

	return create, nil

}
func (s *ServU) GetUser(ctx context.Context, id string) (models.Users, error) {
	return s.SS.GetUser(ctx, id)
}

func (s *ServU) GetUsers(ctx context.Context) ([]models.Users, error) {
	return s.SS.GetUsers(ctx)
}

func (s *ServU) UpdateUsers(ctx context.Context, id string, updateduser *models.Users) (models.Users, error) {
	return s.SS.UpdateUsers(ctx, id, updateduser)
}

func (s *ServU) DeleteUser(ctx context.Context, id string) (string, error) {
	return s.SS.DeleteUser(ctx, id)
}

func (s *ServU) SendEmail(recpiantEmail, fullname string) {
	subject := "Welcome to john jersey shop"
	body := fmt.Sprintf("Hello %s , \nwelcome to john jersey shop we will serve you well !", fullname)

	if err := s.email.Send(recpiantEmail, subject, body); err != nil {
		return
	}
	fmt.Printf("✅ Email Sucessfully sent to %s", fullname)
}

func (s *ServU) EnableAccount(ctx context.Context, id string) (string, error) {
	return s.SS.EnableAccount(ctx, id)
}

func (s *ServU) LoginUser(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	return s.SS.LoginUser(ctx, req)
}
