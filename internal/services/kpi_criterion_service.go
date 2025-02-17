package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPICriterionService struct {
	kpiCriterionRepo *repositories.KPICriterionRepository
}

func NewKPICriterionService(kpiCriterionRepo *repositories.KPICriterionRepository) *KPICriterionService {
	return &KPICriterionService{
		kpiCriterionRepo: kpiCriterionRepo,
	}
}

func (s *KPICriterionService) Create(ctx context.Context, criterion *models.KPICriterion) error {
	return s.kpiCriterionRepo.Create(ctx, criterion)
}

func (s *KPICriterionService) Update(ctx context.Context, criterion *models.KPICriterion) error {
	return s.kpiCriterionRepo.Update(ctx, criterion)
}

func (s *KPICriterionService) Delete(ctx context.Context, id string) error {
	return s.kpiCriterionRepo.Delete(ctx, id)
}

func (s *KPICriterionService) GetByID(ctx context.Context, id string) (*models.KPICriterion, error) {
	return s.kpiCriterionRepo.FindByID(ctx, id)
}

func (s *KPICriterionService) List(ctx context.Context, filter *models.ListKPICriterionRequest) (*models.ListKPICriterionResponse, error) {
	return s.kpiCriterionRepo.FindAll(ctx, filter)
}