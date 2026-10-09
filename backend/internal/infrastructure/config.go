package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv   string
	HTTPPort string
}

func LoadConfig() *Config {
	// 如果没有 .env（例如线上环境），忽略错误即可
	_ = godotenv.Load()
	return &Config{
		AppEnv:   getEnv("APP_ENV", "dev"),
		HTTPPort: getEnv("HTTP_PORT", "8083"),
	}
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
