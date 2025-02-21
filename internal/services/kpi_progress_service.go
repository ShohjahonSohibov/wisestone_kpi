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

func (s *KPIProgressService) Update(ctx context.Context, req *models.KPIProgress) error {
	return s.kpiProgressRepo.Update(ctx, req)
}

func (s *KPIProgressService) Delete(ctx context.Context, id string) error {
	return s.kpiProgressRepo.Delete(ctx, id)
}

func (s *KPIProgressService) GetByID(ctx context.Context, id string) (*models.KPIProgress, error) {
	return s.kpiProgressRepo.GetByID(ctx, id)
}

func (s *KPIProgressService) List(ctx context.Context, req *models.ListKPIProgressRequest) (*models.ListKPIProgressResponse, error) {
	var (
		res *models.ListKPIProgressResponse
		err error
	)
	if req.TeamId != "" {
		filter := models.KPIProgressTeamFilter{
			Limit:  req.Limit,
			Offset: req.Offset,
			Date:   req.Date,
			TeamId: req.TeamId,
		}

		res, err = s.kpiProgressRepo.TeamProgress(ctx, &filter)
		if err != nil {
			return nil, err
		}
	} else if req.EmployeeId != "" {
		// filter := models.KPIProgressEmployeeFilter{
		// 	Limit: req.Limit,
		// 	Offset: req.Offset,
		// 	Date: req.Date,
		// 	EmployeeId: req.EmployeeId,
		// }

		// res, err := s.kpiProgressRepo.List(ctx, &filter)
		// if err != nil {
		// 	return nil, err
		// }
	}
	return res, nil
}
