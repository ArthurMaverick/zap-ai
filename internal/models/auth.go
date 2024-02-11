package models

import (
	"github.com/ArthurMaverick/zap-ai/pkg/bcrypt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EntityUsers struct {
	ID        string `gorm:"primary_key"`
	FullName  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	pw, err := bcrypt.HashPassword(entity.Password)
	if err != nil {
		return err
	}
	entity.Password = pw
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
