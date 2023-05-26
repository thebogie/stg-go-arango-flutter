// internal/service/auth_service.go
package service

import (
	"back/internal/domain"
	repo "back/internal/repository"
)

type AuthService interface {
	LoginUser(email, password string) (*domain.AuthPayload, error)
}

type authService struct {
	authRepository repo.AuthRepository
}

func NewAuthService(authRepository repo.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

// Implement the UserService methods
func (s *authService) LoginUser(email, password string) (*domain.AuthPayload, error) {
	// Implement the logic to register a user
	return nil, nil
}
