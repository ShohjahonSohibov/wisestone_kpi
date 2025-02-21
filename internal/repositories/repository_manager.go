package repositories

import "go.mongodb.org/mongo-driver/mongo"

type RepositoryManager struct {
	UserRepository               *UserRepository
	TeamRepository               *TeamRepository
	RoleRepository               *RoleRepository
	PermissionRepository         *PermissionRepository
	RolePermissionRepository     *RolePermissionRepository
	KPIParentRepository          *KpiParentRepository
	KPICriterionRepository       *KPICriterionRepository
	KPIDivisionRepository        *KPIDivisionRepository
	KPIFactorRepository          *KPIFactorRepository
	KPIFactorIndicatorRepository *KPIFactorIndicatorRepository
	KPIProgressRepository        *KPIProgressRepository
}

func NewRepositoryManager(db *mongo.Database) *RepositoryManager {
	return &RepositoryManager{
		UserRepository:               NewUserRepository(db),
		TeamRepository:               NewTeamRepository(db),
		RoleRepository:               NewRoleRepository(db),
		PermissionRepository:         NewPermissionRepository(db),
		RolePermissionRepository:     NewRolePermissionRepository(db),
		KPIParentRepository:          NewKPIParentRepository(db),
		KPICriterionRepository:       NewKPICriterionRepository(db),
		KPIDivisionRepository:        NewKPIDivisionRepository(db),
		KPIFactorRepository:          NewKPIFactorRepository(db),
		KPIFactorIndicatorRepository: NewKPIFactorIndicatorRepository(db),
		KPIProgressRepository:        NewKPIProgressRepository(db),
	}
}
