package services

import "kpi/internal/repositories"

type ServiceManager struct {
	AuthService           *AuthService
	UserService           *UserService
	TeamService           *TeamService
	RoleService           *RoleService
	PermissionService     *PermissionService
	RolePermissionService *RolePermissionService
}

func NewServiceManager(repoManager *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		AuthService:           NewAuthService(repoManager.UserRepository, repoManager.RoleRepository),
		UserService:           NewUserService(repoManager.UserRepository, repoManager.TeamRepository),
		TeamService:           NewTeamService(repoManager.TeamRepository),
		RoleService:           NewRoleService(repoManager.RoleRepository),
		PermissionService:     NewPermissionService(repoManager.PermissionRepository),
		RolePermissionService: NewRolePermissionService(repoManager.RolePermissionRepository),
	}
}
