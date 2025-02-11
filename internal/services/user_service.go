package services

import (
	"context"
	"errors"
	"kpi/internal/models"
	"kpi/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	existingUser, _ := s.userRepo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	existingUser, err := s.userRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, email string) error {
	existingUser, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}
	return s.userRepo.Delete(ctx, email)
}

func (s *UserService) List(ctx context.Context, filter *models.ListUsersRequest) (*models.ListUsersResponse, error) {
	return s.userRepo.FindAll(ctx, filter)
}
