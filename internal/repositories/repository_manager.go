package repositories

import "go.mongodb.org/mongo-driver/mongo"

type RepositoryManager struct {
    UserRepository *UserRepository
    TeamRepository *TeamRepository
    RoleRepository *RoleRepository
    PermissionRepository *PermissionRepository
    RolePermissionRepository *RolePermissionRepository
    KpiParentRepository *KpiParentRepository
}

func NewRepositoryManager(db *mongo.Database) *RepositoryManager {
    return &RepositoryManager{
        UserRepository: NewUserRepository(db),
        TeamRepository: NewTeamRepository(db),
        RoleRepository: NewRoleRepository(db),
        PermissionRepository: NewPermissionRepository(db),
        RolePermissionRepository: NewRolePermissionRepository(db),
        KpiParentRepository: NewKPIParentRepository(db),
    }
}