package config

import "os"

type Config struct {
	AppName     string
	Version     string
	Environment string
	Port        string
}

func Load() Config {
	return Config{
		AppName:     getEnv("APP_NAME", "DevOps Dashboard"),
		Version:     getEnv("APP_VERSION", "1.0.0"),
		Environment: getEnv("APP_ENV", "Development"),
		Port:        getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
