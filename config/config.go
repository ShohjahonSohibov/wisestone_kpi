package config

import (
	"os"
)

type Config struct {
	DBUri         string
	Port          string
	Secret        string
	DefaultOffset string
	DefaultLimit  string
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Load() *Config {
	return &Config{
		DBUri:         getEnvOrDefault("MONGODB_URI", "mongodb+srv://shohjahon:new@cluster0.qm5mkpn.mongodb.net/"),
		Port:          getEnvOrDefault("SERVER_PORT", "8080"),
		Secret:        getEnvOrDefault("JWT_SECRET", "top"),
		DefaultOffset: getEnvOrDefault("DEFAULT_OFFSET", "0"),
		DefaultLimit:  getEnvOrDefault("DEFAULT_LIMIT", "10"),
	}
}
