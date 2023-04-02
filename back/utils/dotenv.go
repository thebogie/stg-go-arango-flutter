package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	env := make(chan string, 1)

	log.Printf("Starting up env: %v", os.Getenv("GO_ENV"))
	if os.Getenv("GO_ENV") == "prod" {
		godotenv.Load(".env.prod")
		env <- os.Getenv(key)
	} else {
		godotenv.Load(".env.dev")
		env <- os.Getenv(key)
	}

	return <-env
}
