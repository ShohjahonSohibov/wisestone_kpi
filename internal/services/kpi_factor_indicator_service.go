package services

import (
	"context"
	"fmt"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIFactorIndicatorService struct {
	kpiFactorIndicatorRepo *repositories.KPIFactorIndicatorRepository
	kpiFactorRepo          *repositories.KPIFactorRepository
}

func NewKPIFactorIndicatorService(
	kpiFactorIndicatorRepo *repositories.KPIFactorIndicatorRepository,
	kpiFactorRepo *repositories.KPIFactorRepository,
) *KPIFactorIndicatorService {
	return &KPIFactorIndicatorService{
		kpiFactorIndicatorRepo: kpiFactorIndicatorRepo,
		kpiFactorRepo:          kpiFactorRepo,
	}
}

func (s *KPIFactorIndicatorService) Create(ctx context.Context, indicator *models.KPIFactorIndicator) error {
	// Get the factor to check its ratio
	factor, err := s.kpiFactorRepo.FindByID(ctx, indicator.FactorID)
	if err != nil {
		return err
	}

	// Check if progress range exceeds factor's ratio
	if float64(indicator.ProgressRange) > factor.Ratio {
		return fmt.Errorf("progress range cannot exceed factor ratio of %v", factor.Ratio)
	}

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
