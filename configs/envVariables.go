package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvVariable(key string) string {
	//TODO: load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error to load env")
	}

	return os.Getenv(key)
}
