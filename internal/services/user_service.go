package services

import (
	"context"
	"errors"
	"fmt"
	"kpi/internal/models"
	"kpi/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepo *repositories.UserRepository
    teamRepo *repositories.TeamRepository
}

func NewUserService(userRepo *repositories.UserRepository, teamRepo *repositories.TeamRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
        teamRepo: teamRepo,
    }
}

func (s *UserService) GetById(ctx context.Context, id string) (*models.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	existingUser, _ := s.userRepo.FindByID(ctx, user.ID)
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
	existingUser, err := s.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return err
	}
	fmt.Println("\nuser:", user)
	if existingUser == nil {
		return errors.New("user not found")
	}
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	existingUser, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}
	return s.userRepo.Delete(ctx, id)
}

func (s *UserService) List(ctx context.Context, filter *models.ListUsersRequest) (*models.ListUsersResponse, error) {
    if filter.TeamId != "" {
        // Get team information first
        team, err := s.teamRepo.FindByID(ctx, filter.TeamId)
        if err != nil {
            return nil, err
        }
        if team == nil {
            return nil, errors.New("team not found")
        }

        // Get users with the filter
        response, err := s.userRepo.FindAll(ctx, filter)
        if err != nil {
            return nil, err
        }

        // Mark team leader
        for _, user := range response.Items {
            if user.ID == team.LeaderId {
                user.IsTeamLeader = true
            }
        }
        return response, nil
    }

    return s.userRepo.FindAll(ctx, filter)
}

func (s *UserService) AssignTeam(ctx context.Context, userID, teamID string) error {
    existingUser, err := s.userRepo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    if existingUser == nil {
        return errors.New("user not found")
    }
    return s.userRepo.AssignTeam(ctx, userID, teamID)
}

func (s *UserService) RemoveFromTeam(ctx context.Context, userID string) error {
    existingUser, err := s.userRepo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    if existingUser == nil {
        return errors.New("user not found")
    }
    return s.userRepo.RemoveFromTeam(ctx, userID)
}
