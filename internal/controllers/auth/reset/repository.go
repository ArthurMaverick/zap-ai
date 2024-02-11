package reset

import (
	model "github.com/ArthurMaverick/zap-ai/internal/models"
	"github.com/ArthurMaverick/zap-ai/pkg/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	ResetRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResetRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password
	users.Active = input.Active

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)
	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "RESET_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "RESET_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	encryptedPassword, err := bcrypt.HashPassword(input.Password)
	if err != nil {
		errorCode <- "RESET_ENCRYPTED_PASSWORD_400"
		return &users, <-errorCode
	}

	users.Password = encryptedPassword
	users.UpdatedAt = time.Now().Local()

	updateNewPassword := db.Debug().Select("password", "update_at").Where("email = ?", input.Email).Updates(users)
	if updateNewPassword.Error != nil {
		errorCode <- "RESET_UPDATE_PASSWORD_400"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &users, <-errorCode

}
