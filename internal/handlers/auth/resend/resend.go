package resend

import (
	"github.com/ArthurMaverick/zap-ai/internal/controllers/auth/resend"
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
	service resend.Service
}

func NewHandlerResend(service resend.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ResendHandler(ctx *gin.Context) {
	var input resend.InputResend
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response.APIResponse(ctx, "Error input", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	resendResult, errResend := h.service.ResendService(&input)

	switch errResend {
	case "RESEND_NOT_FOUND_404":
		response.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodPost, nil)
		return
	case "RESEND_ALREADY_ACTIVE_400":
		response.APIResponse(ctx, "Email has been active", http.StatusBadRequest, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": resendResult.ID, "email": resendResult.Email}
		accessToken, errToken := jwt.Sign(accessTokenData, "JWT_SECRET", 5)
		if errToken != nil {
			defer slog.Error(errToken.Error())
			response.APIResponse(ctx, "Error generate token", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}

		_, errSendGrid := mail.SendGridMail(resendResult.FullName, resendResult.Email, "Resend new Activation", "template_resend", accessToken)
		if errSendGrid != nil {
			defer slog.Error(errSendGrid.Error())
			response.APIResponse(ctx, "Error send email", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}

		response.APIResponse(ctx, "Success resend activation", http.StatusOK, http.MethodPost, nil)
	}
}
