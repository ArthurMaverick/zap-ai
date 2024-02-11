package env

import (
	"github.com/joho/godotenv"
	"os"
)

// GodoEnv is a function that loads the .env file and returns the value of the key
func GodoEnv(key string) (string, error) {
	env := make(chan string, 1)
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			return "", err
		}
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}
	return <-env, nil
}
