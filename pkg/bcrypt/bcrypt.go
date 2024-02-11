package bcrypt

import (
	"github.com/ArthurMaverick/zap-ai/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	logger.Log()
}

func HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(hasPassword, password string) error {
	pw := []byte(password)
	hw := []byte(hasPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
