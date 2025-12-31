package config

import "os"

type config struct {
	DatabaseUrl string
	Port        string
}

func Load() *config {
	return &config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
