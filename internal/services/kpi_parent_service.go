package services

import (
	"context"

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

func (s *KPIParentService) Delete(ctx context.Context, id string) error {
	return s.kpiParentRepo.Delete(ctx, id)
}

func (s *KPIParentService) GetByID(ctx context.Context, id string) (*models.KPIParent, error) {
	return s.kpiParentRepo.GetByID(ctx, id)
}

func (s *KPIParentService) List(ctx context.Context, req *models.ListKPIParentRequest) (*models.ListKPIParentResponse, error) {
	return s.kpiParentRepo.List(ctx, req)
}
