package services

import (
	"context"
	"fmt"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIParentService struct {
	kpiParentRepo *repositories.KpiParentRepository
}

func NewKPIParentService(kpiParentRepo *repositories.KpiParentRepository) *KPIParentService {
	return &KPIParentService{kpiParentRepo: kpiParentRepo}
}

func (s *KPIParentService) Create(ctx context.Context, req *models.KPIParent) error {
	err := s.kpiParentRepo.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *KPIParentService) Update(ctx context.Context, req *models.KPIParent) error {
	if err := s.kpiParentRepo.Update(ctx, req); err != nil {
		return err
	}

	return nil
}

func (s *KPIParentService) UpdateStatus(ctx context.Context, req *models.UpdateKPIParentStatus) error {
	// Validate status
	switch models.KPIStatus(req.Status) {
	case models.KPIStatusDraft, models.KPIStatusPending, models.KPIStatusApproved,
		models.KPIStatusRejected, models.KPIStatusCancelled:
		// Valid status
	default:
		return fmt.Errorf("invalid status: %s", req.Status)
	}

	return s.kpiParentRepo.UpdateStatus(ctx, req.ID, req.Status)
}

func (s *KPIParentService) Delete(ctx context.Context, id string) error {
	return s.kpiParentRepo.Delete(ctx, id)
}

func (s *KPIParentService) GetByID(ctx context.Context, id, kpiType string) (*models.KPIParent, error) {
	return s.kpiParentRepo.GetByID(ctx, id, kpiType)
}

func (s *KPIParentService) List(ctx context.Context, req *models.ListKPIParentRequest) (*models.ListKPIParentResponse, error) {
	return s.kpiParentRepo.List(ctx, req)
}
