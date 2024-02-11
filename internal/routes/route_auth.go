package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	groupRoute := route.Group("/api/v1")
	{
		groupRoute.GET("/login", func(context *gin.Context) {})
	}
}
