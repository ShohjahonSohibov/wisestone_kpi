package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KpiParentService struct {
	kpiParentRepo *repositories.KpiParentRepository
}

func NewKPIParentService(kpiParentRepo *repositories.KpiParentRepository) *KpiParentService {
	return &KpiParentService{kpiParentRepo: kpiParentRepo}
}

func (s *KpiParentService) Create(ctx context.Context, req *models.KPIParent) error {
	err := s.kpiParentRepo.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *KpiParentService) Update(ctx context.Context, req *models.KPIParent) error {
	if err := s.kpiParentRepo.Update(ctx, req); err != nil {
		return err
	}

	return nil
}

func (s *KpiParentService) Delete(ctx context.Context, id string) error {
	return s.kpiParentRepo.Delete(ctx, id)
}

func (s *KpiParentService) GetByID(ctx context.Context, id string) (*models.KPIParent, error) {
	return s.kpiParentRepo.GetByID(ctx, id)
}

func (s *KpiParentService) List(ctx context.Context, req *models.ListKPIParentRequest) (*models.ListKPIParentResponse, error) {
	return s.kpiParentRepo.List(ctx, req)
}
