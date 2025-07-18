package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("ENV file not found")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
