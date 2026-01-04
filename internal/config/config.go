package config

import "os"

type config struct {
	DatabaseUrl  string
	Port         string
	JWTSecret    string
	JWTExpiryHrs int
}

func Load() *config {
	return &config{
		DatabaseUrl:  os.Getenv("DATABASE_URL"),
		Port:         getEnv("PORT", "8080"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		JWTExpiryHrs: 24,
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
