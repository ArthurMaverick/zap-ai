package main

import (
	"github.com/ArthurMaverick/zap-ai/internal/configs"
	route "github.com/ArthurMaverick/zap-ai/internal/routes"
	"github.com/ArthurMaverick/zap-ai/pkg/env"
	"github.com/ArthurMaverick/zap-ai/pkg/logger"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"os"
)

func init() {
	logger.Log()
}

func main() {

	router := SetupRouter()
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}

func SetupRouter() *gin.Engine {
	db, err := configs.Connection()
	if err != nil {
		slog.Error(err.Error())
	}

	router := gin.Default()

	goEnv, err := env.GodoEnv("GO_ENV")
	if err != nil {
		slog.Error(err.Error())
	}

	if goEnv == "test" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	route.InitAuthRoutes(db, router)
	return router
}
