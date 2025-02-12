package services

import (
	"context"
	"errors"
	"fmt"
	"kpi/internal/models"
	"kpi/internal/repositories"
)

type PermissionService struct {
	permissionRepo *repositories.PermissionRepository
}

func NewPermissionService(permissionRepo *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{permissionRepo: permissionRepo}
}

func (s *PermissionService) GetById(ctx context.Context, id string) (*models.Permission, error) {
	permission, err := s.permissionRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if permission == nil {
		return nil, errors.New("permission not found")
	}
	return permission, nil
}

func (s *PermissionService) Create(ctx context.Context, permission *models.Permission) error {
    // Check if permission with same action_kr already exists
    existing, err := s.permissionRepo.FindByAction(ctx, permission.ActionKr)
    if err != nil {
        return err
    }
    if existing != nil {
        return fmt.Errorf("permission with action_kr '%s' already exists", permission.ActionKr)
    }

    return s.permissionRepo.Create(ctx, permission)
}

func (s *PermissionService) Update(ctx context.Context, permission *models.Permission) error {
	existingPermission, err := s.permissionRepo.FindByID(ctx, permission.ID)
	if err != nil {
		return err
	}
	if existingPermission == nil {
		return errors.New("permission not found")
	}
	return s.permissionRepo.Update(ctx, permission)
}

func (s *PermissionService) Delete(ctx context.Context, id string) error {
	existingPermission, err := s.permissionRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existingPermission == nil {
		return errors.New("permission not found")
	}
	return s.permissionRepo.Delete(ctx, id)
}

func (s *PermissionService) List(ctx context.Context, filter *models.ListPermissionRequest) (*models.ListPermissionResponse, error) {
	return s.permissionRepo.FindAll(ctx, filter)
}
