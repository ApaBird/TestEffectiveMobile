package postgres

import (
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
}

func LoadConfig() (*DBConfig, error) {
	dbConfig := &DBConfig{
		Host:     GetEnv("DB_HOST", "localhost"),
		Port:     GetEnv("DB_PORT", "5432"),
		User:     GetEnv("DB_USER", "postgres"),
		Password: GetEnv("DB_PASSWORD", "postgres"),
		DBName:   GetEnv("DB_NAME", "postgres"),
		Sslmode:  GetEnv("DB_SSLMODE", "disable"),
	}

	return dbConfig, nil
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
