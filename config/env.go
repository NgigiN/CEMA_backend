// This file sets up configs for the application using environment variables or defaults.
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// config struct necessary to load env variables for db connection
type Config struct {
	Host       string `env:"HOST" envDefault:"localhost"`
	Port       string `env:"PORT" envDefault:"8080"`
	DBUSER     string `env:"DB_USER" envDefault:"your_db_user"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"your_db_password"`
	DBAddress  string `env:"DB_ADDRESS" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBName     string `env:"DB_NAME" envDefault:"your_db_name"`
}

var Envs = initConfig()

// initConfig initializes the configuration by loading environment variables
func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:       getEnv("HOST", "localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUSER:     getEnv("DB_USER", "your_db_user"),
		DBPassword: getEnv("DB_PASSWORD", "your_db_password"),
		DBAddress:  getEnv("DB_ADDRESS", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "your_db_name"),
	}
}

// getEnv retrieves the value of an environment variable or returns a fallback value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
