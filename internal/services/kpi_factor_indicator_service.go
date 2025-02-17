package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIFactorIndicatorService struct {
	kpiFactorIndicatorRepo *repositories.KPIFactorIndicatorRepository
}

func NewKPIFactorIndicatorService(kpiFactorIndicatorRepo *repositories.KPIFactorIndicatorRepository) *KPIFactorIndicatorService {
	return &KPIFactorIndicatorService{
		kpiFactorIndicatorRepo: kpiFactorIndicatorRepo,
	}
}

func (s *KPIFactorIndicatorService) Create(ctx context.Context, indicator *models.KPIFactorIndicator) error {
	return s.kpiFactorIndicatorRepo.Create(ctx, indicator)
}

func (s *KPIFactorIndicatorService) Update(ctx context.Context, indicator *models.KPIFactorIndicator) error {
	return s.kpiFactorIndicatorRepo.Update(ctx, indicator)
}

func (s *KPIFactorIndicatorService) Delete(ctx context.Context, id string) error {
	return s.kpiFactorIndicatorRepo.Delete(ctx, id)
}

func (s *KPIFactorIndicatorService) GetByID(ctx context.Context, id string) (*models.KPIFactorIndicator, error) {
	return s.kpiFactorIndicatorRepo.FindByID(ctx, id)
}

func (s *KPIFactorIndicatorService) List(ctx context.Context, filter *models.ListKPIFactorIndicatorRequest) (*models.ListKPIFactorIndicatorResponse, error) {
	return s.kpiFactorIndicatorRepo.FindAll(ctx, filter)
}