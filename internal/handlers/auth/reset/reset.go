package reset

import (
	"github.com/ArthurMaverick/zap-ai/internal/controllers/auth/reset"
	"github.com/ArthurMaverick/zap-ai/pkg/jwt"
	"github.com/ArthurMaverick/zap-ai/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service reset.Service
}

func NewHandlerReset(service reset.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ResetHandler(ctx *gin.Context) {
	var input reset.InputReset
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response.APIResponse(ctx, "Failed to bind input", http.StatusBadRequest, http.MethodConnect, nil)
	}

	token := ctx.Param("token")
	resultToken, errToken := jwt.VerifyToken(token, "JWT_SECRET")

	if errToken != nil {
		response.APIResponse(ctx, "Failed to verify token", http.StatusBadRequest, http.MethodConnect, nil)
		return
	}

	if input.ChangePassword != input.Password {
		response.APIResponse(ctx, "Password and Confirm Password not match", http.StatusBadRequest, http.MethodConnect, nil)
		return
	}

	result := jwt.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, errReset := h.service.ResetService(&input)

	switch errReset {
	case "RESET_NOT_FOUND_404":
		response.APIResponse(ctx, "User account is not exist", http.StatusNotFound, http.MethodPost, nil)
		return
	case "ACCOUNT_NOT_ACTIVE_403":
		response.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return
	case "RESET_PASSWORD_FAILED_403":
		response.APIResponse(ctx, "Change new password failed", http.StatusForbidden, http.MethodPost, nil)
		return
	case "RESET_UPDATE_PASSWORD_400":
		response.APIResponse(ctx, "Change new password failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	default:
		response.APIResponse(ctx, "Change new password successfully", http.StatusOK, http.MethodPost, nil)
	}

}
