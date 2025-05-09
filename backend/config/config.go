package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StoreConf StoreConfig
	DbConf    DbConfig
}

type DbConfig struct {
	DbPort     string
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbUrl      string
}

type StoreConfig struct {
	Port      string
	SecretKey string
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		DbConf: DbConfig{
			DbHost:     os.Getenv("DATABASE_HOST"),
			DbPort:     os.Getenv("DATABASE_PORT"),
			DbUser:     os.Getenv("DATABASE_USER"),
			DbPassword: os.Getenv("DATABASE_PASSWORD"),
			DbName:     os.Getenv("DATABASE_NAME"),
			DbUrl:      os.Getenv("DATABASE_URL"),
		},
		StoreConf: StoreConfig{
			Port:      os.Getenv("PORT"),
			SecretKey: os.Getenv("KEY"),
		},
	}
}
