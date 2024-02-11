package configs

import (
	model "github.com/ArthurMaverick/zap-ai/internal/models"
	"github.com/ArthurMaverick/zap-ai/pkg/env"
	"github.com/ArthurMaverick/zap-ai/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func init() {
	logger.Log()
}

func Connection() (*gorm.DB, error) {
	databaseURI := make(chan string, 1)

	dbEnv, err := env.GodoEnv("DATABASE_URI_DEV")
	if err != nil {
		return nil, err
	}

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- dbEnv
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("GO_ENV") != "production" {
		//db.Debug()
		slog.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
		//&model.EntityStudent{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
