package services

import (
	"context"
	"errors"
	"kpi/internal/models"
	"kpi/internal/repositories"
)

type RolePermissionService struct {
	rolePermissionRepo *repositories.RolePermissionRepository
}

func NewRolePermissionService(rolePermissionRepo *repositories.RolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{rolePermissionRepo: rolePermissionRepo}
}

func (s *RolePermissionService) GetById(ctx context.Context, id string) (*models.RolePermission, error) {
	rolePermission, err := s.rolePermissionRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if rolePermission == nil {
		return nil, errors.New("role permission not found")
	}
	return rolePermission, nil
}

func (s *RolePermissionService) Create(ctx context.Context, rolePermission *models.RolePermission) error {
	return s.rolePermissionRepo.Create(ctx, rolePermission)
}

func (s *RolePermissionService) Update(ctx context.Context, rolePermission *models.RolePermission) error {
	existingRolePermission, err := s.rolePermissionRepo.FindByID(ctx, rolePermission.ID)
	if err != nil {
		return err
	}
	if existingRolePermission == nil {
		return errors.New("role permission not found")
	}
	return s.rolePermissionRepo.Update(ctx, rolePermission)
}

func (s *RolePermissionService) Delete(ctx context.Context, id string) error {
	existingRolePermission, err := s.rolePermissionRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existingRolePermission == nil {
		return errors.New("role permission not found")
	}
	return s.rolePermissionRepo.Delete(ctx, id)
}

func (s *RolePermissionService) List(ctx context.Context, filter *models.ListRolePermissionRequest) (*models.ListRolePermissionResponse, error) {
	return s.rolePermissionRepo.FindAll(ctx, filter)
}