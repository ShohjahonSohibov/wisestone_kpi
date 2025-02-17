package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIDivisionService struct {
	kpiDivisionRepo *repositories.KPIDivisionRepository
}

func NewKPIDivisionService(kpiDivisionRepo *repositories.KPIDivisionRepository) *KPIDivisionService {
	return &KPIDivisionService{
		kpiDivisionRepo: kpiDivisionRepo,
	}
}

func (s *KPIDivisionService) Create(ctx context.Context, division *models.KPIDivision) error {
	return s.kpiDivisionRepo.Create(ctx, division)
}

func (s *KPIDivisionService) Update(ctx context.Context, division *models.KPIDivision) error {
	return s.kpiDivisionRepo.Update(ctx, division)
}

func (s *KPIDivisionService) Delete(ctx context.Context, id string) error {
	return s.kpiDivisionRepo.Delete(ctx, id)
}

func (s *KPIDivisionService) GetByID(ctx context.Context, id string) (*models.KPIDivision, error) {
	return s.kpiDivisionRepo.FindByID(ctx, id)
}

func (s *KPIDivisionService) List(ctx context.Context, filter *models.ListKPIDivisionRequest) (*models.ListKPIDivisionResponse, error) {
	return s.kpiDivisionRepo.FindAll(ctx, filter)
}