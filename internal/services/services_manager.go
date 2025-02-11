package services

import "kpi/internal/repositories"

type ServiceManager struct {
	AuthService *AuthService
	UserService *UserService
}

func NewServiceManager(repoManager *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		AuthService: NewAuthService(repoManager.UserRepository),
		UserService: NewUserService(repoManager.UserRepository),
	}
}