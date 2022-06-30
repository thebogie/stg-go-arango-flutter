package route

import (
	config "back/configs"
	activationAuth "back/controllers/auth-controllers/activation"
	forgotAuth "back/controllers/auth-controllers/forgot"
	loginAuth "back/controllers/auth-controllers/login"
	registerAuth "back/controllers/auth-controllers/register"
	resendAuth "back/controllers/auth-controllers/resend"
	resetAuth "back/controllers/auth-controllers/reset"
	handlerActivation "back/handlers/auth-handlers/activation"
	handlerForgot "back/handlers/auth-handlers/forgot"
	handlerLogin "back/handlers/auth-handlers/login"
	handlerRegister "back/handlers/auth-handlers/register"
	handlerResend "back/handlers/auth-handlers/resend"
	handlerReset "back/handlers/auth-handlers/reset"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(db *config.DBconnection, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	activationRepository := activationAuth.NewRepositoryActivation(db)
	activationService := activationAuth.NewServiceActivation(activationRepository)
	activationHandler := handlerActivation.NewHandlerActivation(activationService)

	resendRepository := resendAuth.NewRepositoryResend(db)
	resendService := resendAuth.NewServiceResend(resendRepository)
	resendHandler := handlerResend.NewHandlerResend(resendService)

	forgotRepository := forgotAuth.NewRepositoryForgot(db)
	forgotService := forgotAuth.NewServiceForgot(forgotRepository)
	forgotHandler := handlerForgot.NewHandlerForgot(forgotService)

	resetRepository := resetAuth.NewRepositoryReset(db)
	resetService := resetAuth.NewServiceReset(resetRepository)
	resetHandler := handlerReset.NewHandlerReset(resetService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")

	groupRoute.POST("/login", loginHandler.LoginHandler)

	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/resend-token", resendHandler.ResendHandler)
	groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
	groupRoute.POST("/change-password/:token", resetHandler.ResetHandler)

}
