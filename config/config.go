package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port           string `env:"PORT" envDefault:"8080"`
	SecretPassword string `env:"SECRET_PASSWORD"`
}

func NewConfig() *Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Using environment variables.")
	}

	config := &Config{}

	// Parse environment variables into the config struct
	if err := env.Parse(config); err != nil {
		log.Fatalln("Failed to parse environment variables into Config struct:", err)
	}

	return config
}
