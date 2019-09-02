package config

import (
	"os"
	"strconv"
)

// DBConfig contains all relevant information to connect to the database
type DBConfig struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
}

// Config contains a DBConfig for database access
type Config struct {
	DB DBConfig
}

// New creates a new instance of a configuration
func New() *Config {
	return &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnvAsUInt16("DB_PORT", 3306),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "root"),
			Database: getEnv("DB_DATABASE", "none"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsUInt16(key string, defaultVal uint16) uint16 {
	valStr := getEnv(key, "")
	if value, err := strconv.Atoi(valStr); err != nil {
		return uint16(value)
	}

	return defaultVal
}
