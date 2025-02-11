package config

import (
	"os"
)

type Config struct {
	DBUri    string
	Port     string
	Secret   string
	DefaultOffset string
	DefaultLimit string
}

func Load() *Config {
	return &Config{
		DBUri:  os.Getenv("MONGODB_URI"),
		Port:   os.Getenv("SERVER_PORT"),
		Secret: os.Getenv("JWT_SECRET"),
		DefaultOffset: os.Getenv("DEFAULT_OFFSET"),
		DefaultLimit: os.Getenv("DEFAULT_LIMIT"),
	}
}
