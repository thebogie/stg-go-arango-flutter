// internal/usecase/user_usecase.go
package usecase

import (
	"back/internal/domain"
	service "back/internal/service"
)

type UserUsecase interface {
	RegisterUser(username, email, password string) (*domain.AuthPayload, error)
}

type userUsecase struct {
	userService service.UserService
}

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{
		userService: userService,
	}
}

// Implement the UserUsecase methods
func (u *userUsecase) RegisterUser(username, email, password string) (*domain.AuthPayload, error) {
	// Implement the logic to register a user
	return nil, nil
}
