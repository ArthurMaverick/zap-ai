package resend

import (
	model "github.com/ArthurMaverick/zap-ai/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	ResendRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewResendRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ResendRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email

	checkUserAccount := db.Debug().Select("*").Where("email = ?", users.Email).Find(&users)
	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "RESEND_NOT_FOUND_404"
		return nil, <-errorCode
	}

	if users.Active {
		errorCode <- "RESEND_ALREADY_ACTIVE_400"
		return &users, <-errorCode
	} else {
		users.Active = true
		db.Save(&users)
		errorCode <- "nil"
	}
	return &users, <-errorCode
}
