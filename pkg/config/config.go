package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadAllConfigs(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}

	LoadAppConfig()
	LoadNocoDBConfig()
	LoadNotionConfig()
}
