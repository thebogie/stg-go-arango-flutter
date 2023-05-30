// internal/usecase/user_usecase.go
package usecase

import (
	"back/internal/domain"
	service "back/internal/service"
	"github.com/sirupsen/logrus"
)

type UserUsecase interface {
	RegisterUser(username, email, password string) (*domain.AuthPayload, error)
	FindUserByEmail(email string) (*domain.User, error)
	FindUserByID(id string) (*domain.User, error)
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

func (u *userUsecase) FindUserByEmail(email string) (*domain.User, error) {
	user, err := u.userService.FindUserByEmail(email)
	logrus.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}

func (u *userUsecase) FindUserByID(userid string) (*domain.User, error) {
	user, err := u.userService.FindUserByID(userid)
	logrus.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}
