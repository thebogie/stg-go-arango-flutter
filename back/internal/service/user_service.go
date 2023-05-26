// internal/service/user_service.go
package service

import (
	"back/internal/domain"
	repo "back/internal/repository"
)

type UserService interface {
	RegisterUser(username, email, password string) (*domain.AuthPayload, error)
}

type userService struct {
	userRepository repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Implement the UserService methods
func (s *userService) RegisterUser(username, email, password string) (*domain.AuthPayload, error) {
	// Implement the logic to register a user
	return nil, nil
}
