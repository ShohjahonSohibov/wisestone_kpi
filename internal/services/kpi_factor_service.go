package services

import (
	"context"
	"fmt"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIFactorService struct {
	kpiFactorRepo    *repositories.KPIFactorRepository
	kpiCriterionRepo *repositories.KPICriterionRepository
}

func NewKPIFactorService(
	kpiFactorRepo *repositories.KPIFactorRepository,
	kpiCriterionRepo *repositories.KPICriterionRepository,
) *KPIFactorService {
	return &KPIFactorService{
		kpiFactorRepo:    kpiFactorRepo,
		kpiCriterionRepo: kpiCriterionRepo,
	}
}

func (s *KPIFactorService) Create(ctx context.Context, factor *models.KPIFactor) error {
	// Get current sum of ratios for this criterion
	currentSum, err := s.kpiFactorRepo.GetSumRatioByCriterionID(ctx, factor.CriterionID)
	if err != nil {
		return err
	}

	// Get the criterion to check its total ratio
	criterion, err := s.kpiCriterionRepo.FindByID(ctx, factor.CriterionID)
	if err != nil {
		return err
	}

	// Check if adding this factor's ratio would exceed the criterion's total ratio
	if float64(currentSum)+factor.Ratio > criterion.TotalRatio {
		return fmt.Errorf("factor ratio cannot exceed criterion ratio of %v", criterion.TotalRatio)
	}

	return s.kpiFactorRepo.Create(ctx, factor)
}

func (s *KPIFactorService) Update(ctx context.Context, factor *models.KPIFactor) error {
	return s.kpiFactorRepo.Update(ctx, factor)
}

func (s *KPIFactorService) Delete(ctx context.Context, id string) error {
	return s.kpiFactorRepo.Delete(ctx, id)
}

func (s *KPIFactorService) GetByID(ctx context.Context, id string) (*models.KPIFactor, error) {
	return s.kpiFactorRepo.FindByID(ctx, id)
}

func (s *KPIFactorService) List(ctx context.Context, filter *models.ListKPIFactorRequest) (*models.ListKPIFactorResponse, error) {
	return s.kpiFactorRepo.FindAll(ctx, filter)
}
