package services

import "kpi/internal/repositories"

type ServiceManager struct {
	AuthService       *AuthService
	UserService       *UserService
	TeamService       *TeamService
	RoleService       *RoleService
	PermissionService *PermissionService
}

func NewServiceManager(repoManager *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		AuthService:       NewAuthService(repoManager.UserRepository, repoManager.RoleRepository),
		UserService:       NewUserService(repoManager.UserRepository),
		TeamService:       NewTeamService(repoManager.TeamRepository),
		RoleService:       NewRoleService(repoManager.RoleRepository),
		PermissionService: NewPermissionService(repoManager.PermissionRepository),
	}
}
