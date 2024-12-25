package config

import (
	"os"

	"github.com/joho/godotenv"
)

func New() Config {
	godotenv.Load()
	return Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
		Token:       os.Getenv("TOKEN"),
	}
}
