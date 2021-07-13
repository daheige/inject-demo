package service

import (
	"log"

	"github.com/daheige/inject-demo/internal/config"
	"github.com/daheige/inject-demo/internal/domain/repo"
)

// UserService user service
type UserService struct {
	config     *config.AppConfig
	repository repo.UserRepository
}

// NewUserService return user service
func NewUserService(config *config.AppConfig, userRepo repo.UserRepository) *UserService {
	return &UserService{
		config:     config,
		repository: userRepo,
	}
}

// FindUsers return users
func (s *UserService) FindUsers() ([]repo.User, error) {
	users, err := s.repository.FindAll()
	if s.config.AppDebug && err != nil {
		log.Println("get users error: ", err)
	}

	return users, err
}
