package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func NewConfig(dbUser, dbPass, dbName string) Config {
	return Config{
		DB_USER: dbUser,
		DB_PASS: dbPass,
		DB_NAME: dbName,
	}
}

func GetConfig() Config {
	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	cfg := NewConfig(dbUser, dbPass, dbName)

	return cfg
}
