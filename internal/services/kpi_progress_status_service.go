package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIProgressStatusService struct {
	repo *repositories.KPIProgressStatusRepository
}

func NewKPIProgressStatusService(repo *repositories.KPIProgressStatusRepository) *KPIProgressStatusService {
	return &KPIProgressStatusService{repo: repo}
}

func (s *KPIProgressStatusService) Create(ctx context.Context, req *models.CreateKPIProgressStatus) error {
	status := &models.KPIProgressStatus{
		TeamId:     req.TeamId,
		EmployeeId: req.EmployeeId,
		Date:       req.Date,
		Status:     req.Status,
	}
	return s.repo.Create(ctx, status)
}

func (s *KPIProgressStatusService) Update(ctx context.Context, req *models.UpdateKPIProgressStatus) error {
	return s.repo.Update(ctx, req)
}

func (s *KPIProgressStatusService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *KPIProgressStatusService) List(ctx context.Context, req *models.ListKPIProgressStatusRequest) (*models.ListKPIProgressStatusResponse, error) {
	return s.repo.List(ctx, req)
}
