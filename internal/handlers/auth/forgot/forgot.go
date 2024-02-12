package forgot

import (
	"github.com/ArthurMaverick/zap-ai/internal/controllers/auth/forgot"
	"github.com/ArthurMaverick/zap-ai/pkg/jwt"
	"github.com/ArthurMaverick/zap-ai/pkg/mail"
	"github.com/ArthurMaverick/zap-ai/pkg/response"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type handler struct {
	service forgot.Service
}

func NewHandlerForgot(service forgot.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ForgotHandler(ctx *gin.Context) {
	var input forgot.InputForgot
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response.APIResponse(ctx, "Failed", 400, "Invalid input", nil)
		return
	}

	// VALIDATE INPUT
	forgotResult, errForgot := h.service.ForgotService(&input)

	switch errForgot {
	case "FORGOT_NOT_FOUND_404":
		response.APIResponse(ctx, "Email is not never registered", http.StatusNotFound, http.MethodPost, nil)
		return
	case "FORGOT_NOT_ACTIVE_403":
		response.APIResponse(ctx, "Email is not active", http.StatusForbidden, http.MethodPost, nil)
		return
	case "FORGOT_PASSWORD_FAILED_403":
		response.APIResponse(ctx, "Forgot password failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": forgotResult.ID, "email": forgotResult.Email}
		accessToken, errToken := jwt.Sign(accessTokenData, "JWT_SECRET", 5)

		if errToken != nil {
			defer slog.Error(errToken.Error())
			response.APIResponse(ctx, "Generate accessToken failed", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}

		_, errMail := mail.SendGridMail(forgotResult.FullName, forgotResult.Email, "Reset Password", "template_reset", accessToken)
		if errMail != nil {
			response.APIResponse(ctx, "Sending email reset password failed", http.StatusInternalServerError, http.MethodPost, nil)
			return
		}

		response.APIResponse(ctx, "Success", http.StatusOK, http.MethodPost, nil)
	}
}
