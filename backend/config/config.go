package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBUrl  string
	AppEnv string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:   getEnv("PORT", "8080"),
		DBUrl:  getEnv("DATABASE_URL", ""),
		AppEnv: getEnv("APP_ENV", "development"),
	}

	if cfg.DBUrl == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
