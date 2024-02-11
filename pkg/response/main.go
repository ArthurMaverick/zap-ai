package response

import (
	"github.com/ArthurMaverick/zap-ai/pkg/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func init() {
	logger.Log()
}

// Responses is a struct for response
type Responses struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// ErrorResponse is a struct for error response
type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

// APIResponse is a function to handle response
func APIResponse(ctx gin.Context, Data interface{}, StatusCode int, Message, Method string) {

	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		slog.Error(jsonResponse.Message)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

// ValidatorErrorResponse is a function to handle error response
func ValidatorErrorResponse(ctx gin.Context, StatusCode int, Method string, Error interface{}) {

	ctx.JSON(StatusCode, ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	})
	defer ctx.AbortWithStatus(StatusCode)
}
