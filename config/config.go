package config

import (
	"os"
	"strconv"
)

type Config struct {
	App struct {
		Port int
	}
	DB struct {
		Host    string
		Port    int
		User    string
		Password string
		Name    string
	}
}

func LoadConfig() (*Config, error) {
	c := &Config{}

    // 環境変数から読み込む例
    c.App.Port = getEnvAsInt("PORT", 8080)
    c.DB.Host = getEnv("DB_HOST", "db")
    c.DB.Port = getEnvAsInt("DB_PORT", 3306)
    c.DB.User = getEnv("DB_USER", "todouser")
    c.DB.Password = getEnv("DB_PASSWORD", "secret")
    c.DB.Name = getEnv("DB_NAME", "todo")

	return c, nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return i
}