package routes

import (
	activationAuth "github.com/ArthurMaverick/zap-ai/internal/controllers/auth/activation"
	forgotAuth "github.com/ArthurMaverick/zap-ai/internal/controllers/auth/forgot"
	loginAuth "github.com/ArthurMaverick/zap-ai/internal/controllers/auth/login"
	registerAuth "github.com/ArthurMaverick/zap-ai/internal/controllers/auth/register"
	resendAuth "github.com/ArthurMaverick/zap-ai/internal/controllers/auth/resend"
	resetAuth "github.com/ArthurMaverick/zap-ai/internal/controllers/auth/reset"
	activationHandlerAuth "github.com/ArthurMaverick/zap-ai/internal/handlers/auth/activation"
	forgotHandlerAuth "github.com/ArthurMaverick/zap-ai/internal/handlers/auth/forgot"
	loginHandlerAuth "github.com/ArthurMaverick/zap-ai/internal/handlers/auth/login"
	registerHandlerAuth "github.com/ArthurMaverick/zap-ai/internal/handlers/auth/register"
	resendHandlerAuth "github.com/ArthurMaverick/zap-ai/internal/handlers/auth/resend"
	resetHandlerAuth "github.com/ArthurMaverick/zap-ai/internal/handlers/auth/reset"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	loginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(loginRepository)
	loginHandler := loginHandlerAuth.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := registerHandlerAuth.NewHandlerRegister(registerService)

	activationRepository := activationAuth.NewRepositoryActivation(db)
	activationService := activationAuth.NewServiceActivation(activationRepository)
	activationHandler := activationHandlerAuth.NewHandlerActivation(activationService)

	resendRepository := resendAuth.NewResendRepository(db)
	resendService := resendAuth.NewServiceResend(resendRepository)
	resendHandler := resendHandlerAuth.NewHandlerResend(resendService)

	forgotRepository := forgotAuth.NewRepository(db)
	forgotService := forgotAuth.NewServiceForgot(forgotRepository)
	forgotHandler := forgotHandlerAuth.NewHandlerForgot(forgotService)

	resetRepository := resetAuth.NewRepository(db)
	resetService := resetAuth.NewService(resetRepository)
	resetHandler := resetHandlerAuth.NewHandlerReset(resetService)

	groupRoute := route.Group("/api/v1")
	{
		groupRoute.POST("/login", loginHandler.LoginHandler)
		groupRoute.POST("/register", registerHandler.RegisterHandler)
		groupRoute.POST("activation/:token", activationHandler.ActivationHandler)
		groupRoute.POST("/resend-token", resendHandler.ResendHandler)
		groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
		groupRoute.POST("/reset-password/:token", resetHandler.ResetHandler)

	}
}
