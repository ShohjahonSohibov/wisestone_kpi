package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIFactorService struct {
	kpiFactorRepo *repositories.KPIFactorRepository
}

func NewKPIFactorService(kpiFactorRepo *repositories.KPIFactorRepository) *KPIFactorService {
	return &KPIFactorService{
		kpiFactorRepo: kpiFactorRepo,
	}
}

func (s *KPIFactorService) Create(ctx context.Context, factor *models.KPIFactor) error {
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