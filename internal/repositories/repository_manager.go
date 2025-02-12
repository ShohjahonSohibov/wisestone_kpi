package repositories

import "go.mongodb.org/mongo-driver/mongo"

type RepositoryManager struct {
	UserRepository *UserRepository
	TeamRepository *TeamRepository
}

func NewRepositoryManager(db *mongo.Database) *RepositoryManager {
	return &RepositoryManager{
		UserRepository: NewUserRepository(db),
		TeamRepository: NewTeamRepository(db),
	}
}