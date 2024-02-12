package login

import (
	"github.com/ArthurMaverick/zap-ai/internal/controllers/auth/login"
	"github.com/ArthurMaverick/zap-ai/pkg/jwt"
	"github.com/ArthurMaverick/zap-ai/pkg/logger"
	"github.com/ArthurMaverick/zap-ai/pkg/response"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func init() {
	logger.Log()
}

type handler struct {
	service login.Service
}

func NewHandlerLogin(service login.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input login.InputLogin
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {
	case "LOGIN_NOT_FOUND_404":
		response.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return
	case "LOGIN_NOT_ACTIVE_403":
		response.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return
	case "LOGIN_WRONG_PASSWORD_403":
		response.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := jwt.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			defer slog.Error(errToken.Error())
			response.APIResponse(ctx, "Failed to generate token", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}
		response.APIResponse(ctx, "Login success", http.StatusOK, http.MethodPost, map[string]interface{}{"access_token": accessToken})
	}

}
