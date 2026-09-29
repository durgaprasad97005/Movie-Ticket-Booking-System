package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	MongoUri string
	DbName string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}

	return &Config{
		Port: os.Getenv("PORT"),
		MongoUri: os.Getenv("MONGO_URI"),
		DbName: os.Getenv("DB_NAME"),
	}
}