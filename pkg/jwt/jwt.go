package jwt

import (
	"encoding/json"
	"github.com/ArthurMaverick/zap-ai/pkg/env"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type MetaToken struct {
	ID            string
	Email         string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecretPublicKeyEnvName string, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()

	jwtSecretKey, err := env.GodoEnv(SecretPublicKeyEnvName)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	jwtSecretKey, err := env.GodoEnv(SecretPublicKeyEnvName)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func VerifyToken(accessToken, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey, err := env.GodoEnv(SecretPublicKeyEnvName)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
