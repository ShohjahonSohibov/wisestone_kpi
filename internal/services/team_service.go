package services

import (
	"context"
	"errors"
	"kpi/internal/models"
	"kpi/internal/repositories"
)

type TeamService struct {
	teamRepo *repositories.TeamRepository
}

func NewTeamService(teamRepo *repositories.TeamRepository) *TeamService {
	return &TeamService{teamRepo: teamRepo}
}

func (s *TeamService) GetById(ctx context.Context, id string) (*models.Team, error) {
	team, err := s.teamRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if team == nil {
		return nil, errors.New("team not found")
	}
	return team, nil
}

func (s *TeamService) Create(ctx context.Context, team *models.Team) error {
	return s.teamRepo.Create(ctx, team)
}

func (s *TeamService) Update(ctx context.Context, team *models.Team) error {
	existingTeam, err := s.teamRepo.FindByID(ctx, team.ID)
	if err != nil {
		return err
	}
	if existingTeam == nil {
		return errors.New("team not found")
	}
	return s.teamRepo.Update(ctx, team)
}

func (s *TeamService) Delete(ctx context.Context, id string) error {
	existingTeam, err := s.teamRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existingTeam == nil {
		return errors.New("team not found")
	}
	return s.teamRepo.Delete(ctx, id)
}

func (s *TeamService) List(ctx context.Context, filter *models.ListTeamsRequest) (*models.ListTeamsResponse, error) {
	return s.teamRepo.FindAll(ctx, filter)
}
