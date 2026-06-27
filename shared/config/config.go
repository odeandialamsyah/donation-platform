package config

import "os"

type Config struct {
	AppEnv string

	MySQLHost     string
	MySQLPort     string
	MySQLDatabase string
	MySQLUser     string
	MySQLPassword string

	RedisHost string
	RedisPort string

	RabbitMQHost string
	RabbitMQPort string
}

func Load() *Config {
	return &Config{
		AppEnv: os.Getenv("APP_ENV"),

		MySQLHost:     os.Getenv("MYSQL_HOST"),
		MySQLPort:     os.Getenv("MYSQL_PORT"),
		MySQLDatabase: os.Getenv("MYSQL_DATABASE"),
		MySQLUser:     os.Getenv("MYSQL_USERNAME"),
		MySQLPassword: os.Getenv("MYSQL_PASSWORD"),

		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),

		RabbitMQHost: os.Getenv("RABBITMQ_HOST"),
		RabbitMQPort: os.Getenv("RABBITMQ_PORT"),
	}
}
