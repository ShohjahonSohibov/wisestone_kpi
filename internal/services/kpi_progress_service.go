package services

import (
	"context"

	"kpi/internal/models"
	"kpi/internal/repositories"
)

type KPIProgressService struct {
	kpiProgressRepo *repositories.KPIProgressRepository
}

func NewKPIProgressService(kpiProgressRepo *repositories.KPIProgressRepository) *KPIProgressService {
	return &KPIProgressService{
		kpiProgressRepo: kpiProgressRepo,
	}
}

func (s *KPIProgressService) Create(ctx context.Context, req *models.KPIProgress) error {
	return s.kpiProgressRepo.Create(ctx, req)
}

func (s *KPIProgressService) CreateMany(ctx context.Context, progresses []*models.CreateBulkKPIProgress) error {
    return s.kpiProgressRepo.CreateMany(ctx, progresses)
}

func (s *KPIProgressService) Delete(ctx context.Context, date, teamId, employeeId string) error {
	return s.kpiProgressRepo.Delete(ctx, date, teamId, employeeId)
}

func (s *KPIProgressService) List(ctx context.Context, req *models.ListKPIProgressRequest) (*models.ListKPIProgressResponse, error) {
	var (
		res *models.ListKPIProgressResponse
		err error
	)
	if req.TeamId != "" {
		filter := models.KPIProgressTeamFilter{
			Date:   req.Date,
			TeamId: req.TeamId,
		}

		res, err = s.kpiProgressRepo.TeamProgress(ctx, &filter)
		if err != nil {
			return nil, err
		}
	} else if req.EmployeeId != "" {
		filter := models.KPIProgressEmployeeFilter{
			Date:       req.Date,
			EmployeeId: req.EmployeeId,
		}

		res, err = s.kpiProgressRepo.EmployeeProgress(ctx, &filter)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
