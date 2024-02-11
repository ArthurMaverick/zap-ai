package forgot

import (
	model "github.com/ArthurMaverick/zap-ai/internal/models"
	"github.com/ArthurMaverick/zap-ai/pkg/bcrypt"
	"github.com/ArthurMaverick/zap-ai/pkg/random"
	"gorm.io/gorm"
)

type Repository interface {
	ForgotRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ForgotRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var users model.EntityUsers

	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	pw, err := bcrypt.HashPassword(random.String(20))
	if err != nil {
		return nil, err.Error()
	}
	users.Password = pw

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)
	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "FORGOT_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "FORGOT_NOT_ACTIVE_400"
		return &users, <-errorCode
	}

	changePassword := db.Debug().Select("password", "updated_at").Where("email = ?", input.Email).Updates(users)
	if changePassword.Error != nil {
		errorCode <- "FORGOT_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &users, <-errorCode
}
