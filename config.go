package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Database string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load env")
	}
	return &Config{
		Port:     getEnv("PORT", ""),
		Database: getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, defaultKey string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultKey

}
