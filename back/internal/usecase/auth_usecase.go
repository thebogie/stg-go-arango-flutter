// internal/usecase/auth_usecase.go
package usecase

import (
	"back/internal/domain"
	service "back/internal/service"
)

type AuthUsecase interface {
	LoginUser(email, password string) (*domain.AuthPayload, error)
}

type authUsecase struct {
	authService service.AuthService
}

func NewAuthUsecase(authService service.AuthService) AuthUsecase {
	return &authUsecase{
		authService: authService,
	}
}

// Implement the authUsecase methods
func (u *authUsecase) LoginUser(email, password string) (*domain.AuthPayload, error) {
	// Implement the logic to register a user
	return nil, nil
}
