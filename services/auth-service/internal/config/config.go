package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	AppPort string

	MySQLHost     string
	MySQLPort     string
	MySQLDatabase string
	MySQLUsername string
	MySQLPassword string

	JWTSecret string
}

func Load() (*Config, error) {

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env not found, using system env")
	}

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),

		MySQLHost:     os.Getenv("MYSQL_HOST"),
		MySQLPort:     os.Getenv("MYSQL_PORT"),
		MySQLDatabase: os.Getenv("MYSQL_DATABASE"),
		MySQLUsername: os.Getenv("MYSQL_USERNAME"),
		MySQLPassword: os.Getenv("MYSQL_PASSWORD"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	// DEBUG penting
	log.Println("Config loaded, APP_PORT =", cfg.AppPort)

	return cfg, nil
}
