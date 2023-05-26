// cmd/main.go
package main

import (
	"log"
	"net/http"

	"back/internal/repository"
	"back/internal/service"
	"back/internal/usecase"
	"back/web"
)

func main() {
	userRepository := repository.NewUserRepository() // Implement your user repository
	authRepository := repository.NewAuthRepository() // Implement your authentication repository

	userService := service.NewUserService(userRepository) // Implement your user service
	authService := service.NewAuthService(authRepository) // Implement your authentication service

	userUsecase := usecase.NewUserUsecase(userService) // Implement your user use case
	authUsecase := usecase.NewAuthUsecase(authService) // Implement your authentication use case

	server := web.NewServer(userUsecase, authUsecase)
	log.Fatal(http.ListenAndServe(":50002", server))
}
