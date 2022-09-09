package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	os.Setenv("GO_ENV", "test")

	argLength := len(os.Args[1:])
	if argLength != 0 && os.Args[1] == "prod" {
		os.Setenv("GO_ENV", "production")
	}

	env := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	} else {
		godotenv.Load(".env.prod")
		env <- os.Getenv(key)
	}

	return <-env
}
