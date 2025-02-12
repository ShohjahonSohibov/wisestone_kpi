package services

import (
	"context"
	"errors"
	"kpi/internal/models"
	"kpi/internal/repositories"
)

type RoleService struct {
	roleRepo *repositories.RoleRepository
}

func NewRoleService(roleRepo *repositories.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) GetById(ctx context.Context, id string) (*models.Role, error) {
	role, err := s.roleRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role not found")
	}
	return role, nil
}

func (s *RoleService) Create(ctx context.Context, role *models.Role) error {
	return s.roleRepo.Create(ctx, role)
}

func (s *RoleService) Update(ctx context.Context, role *models.Role) error {
	existingRole, err := s.roleRepo.FindByID(ctx, role.ID)
	if err != nil {
		return err
	}
	if existingRole == nil {
		return errors.New("role not found")
	}
	return s.roleRepo.Update(ctx, role)
}

func (s *RoleService) Delete(ctx context.Context, id string) error {
	existingRole, err := s.roleRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existingRole == nil {
		return errors.New("role not found")
	}
	return s.roleRepo.Delete(ctx, id)
}

func (s *RoleService) List(ctx context.Context, filter *models.ListRoleRequest) (*models.ListRoleResponse, error) {
	return s.roleRepo.FindAll(ctx, filter)
}