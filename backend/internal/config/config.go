package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbUri      string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	config :=
		Config{
			Port:       getEnv("PORT", "8080"),
			DbPort:     getEnv("DB_PORT", "27017"),
			DbHost:     getEnv("DB_HOST", "localhost"),
			DbUser:     getEnv("DB_USER", "user"),
			DbPassword: getEnv("DB_PASSWORD", "password"),
			DbName:     getEnv("DB_NAME", "dbname"),
			DbUri:      "",
		}

	config.DbUri = "mongodb://" + config.DbHost + ":" + config.DbPort + "/" + config.DbName

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
