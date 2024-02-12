package register

import (
	"github.com/ArthurMaverick/zap-ai/internal/controllers/auth/register"
	"github.com/ArthurMaverick/zap-ai/pkg/env"
	"github.com/ArthurMaverick/zap-ai/pkg/jwt"
	"github.com/ArthurMaverick/zap-ai/pkg/logger"
	"github.com/ArthurMaverick/zap-ai/pkg/mail"
	"github.com/ArthurMaverick/zap-ai/pkg/response"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func init() {
	logger.Log()
}

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.InputRegister
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.APIResponse(ctx, "Error input", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	resultRegister, errRegister := h.service.RegisterService(&input)

	switch errRegister {
	case "REGISTER_CONFLICT_409":
		response.APIResponse(ctx, "Email already registered", http.StatusConflict, http.MethodPost, nil)
		return
	case "REGISTER_FAILED_403":
		response.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		// TODO : handle this exception other layer
		jwtSecret, err := env.GodoEnv("JWT_SECRET")
		if err != nil {
			slog.Error("Error get JWT_SECRET from env file")
		}
		accessToken, errToken := jwt.Sign(accessTokenData, jwtSecret, 60)
		if errToken != nil {
			defer slog.Error(errToken.Error())
			response.APIResponse(ctx, "Generate accessToken failed", http.StatusForbidden, http.MethodPost, nil)
		}

		_, errSendMail := mail.SendGridMail(resultRegister.FullName, resultRegister.Email, "Activation Account", "template_register", accessToken)
		if errSendMail != nil {
			defer slog.Error(errSendMail.Error())
			response.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}
		response.APIResponse(ctx, "Register new account success", http.StatusCreated, http.MethodPost, nil)
	}
}
