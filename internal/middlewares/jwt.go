package middlewares

import (
	"github.com/ArthurMaverick/zap-ai/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UnauthorizedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var errorResponse UnauthorizedError

		errorResponse.Status = "Forbidden"
		errorResponse.Code = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "You are not authorized to access this resource"

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errorResponse)
			defer ctx.AbortWithStatus(http.StatusForbidden)
		}

		token, err := jwt.VerifyTokenHeader(ctx, "JWT_SECRET")

		errorResponse.Status = "Unauthorized"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "access token invalid or expired"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			ctx.Set("user", token.Claims)
			// return to next method if token is existed
			ctx.Next()
		}
	})
}
