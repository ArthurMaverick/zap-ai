package activation

import (
	"github.com/ArthurMaverick/zap-ai/internal/controllers/auth/activation"
	"github.com/ArthurMaverick/zap-ai/pkg/jwt"
	"github.com/ArthurMaverick/zap-ai/pkg/response"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type handler struct {
	service activation.Service
}

func NewHandler(service activation.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ActivationHandler(ctx *gin.Context) {
	var input activation.InputActivation

	// HANDLER INPUT

	token := ctx.Param("token")
	resultToken, errToken := jwt.VerifyToken(token, "JWT_SECRET")

	if errToken != nil {
		defer slog.Error(errToken.Error())
		response.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := jwt.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, errActivation := h.service.ActivationService(&input)

	switch errActivation {
	case "ACTIVATION_NOT_FOUND_404":
		response.APIResponse(ctx, "User account is not exist", http.StatusNotFound, http.MethodPost, nil)
		return
	case "ACTIVATION_ACTIVE_400":
		response.APIResponse(ctx, "User account is already active", http.StatusBadRequest, http.MethodPost, nil)
		return
	case "ACTIVATION_ACCOUNT_FAILED_403":
		response.APIResponse(ctx, "User account activation failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		response.APIResponse(ctx, "User account activation success", http.StatusOK, http.MethodPost, nil)
	}
}
