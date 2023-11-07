package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Scope                     string
	CalendarCredentialsBase64 string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Scope:                     os.Getenv("GOOGLE_CAL_SCOPE"),
		CalendarCredentialsBase64: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_BASE64"),
	}
}
