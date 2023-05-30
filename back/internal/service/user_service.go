// internal/service/user_service.go
package service

import (
	"back/internal/domain"
	repo "back/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(username, email, password string) (*domain.AuthPayload, error)
	FindUserByEmail(email string) (*domain.User, error)
	FindUserByID(userid string) (*domain.User, error)
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

func (s *userService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	logrus.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}

func (s *userService) FindUserByID(id string) (*domain.User, error) {
	user, err := s.userRepository.FindUserByID(id)
	logrus.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}
