package config

import (
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/joho/godotenv"
)

func LoadAllConfigs(envFile string) {
	log := logger.GetLogger()

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}

	LoadAppConfig()
	LoadNocoDBConfig()
	LoadNotionConfig()

	log.Info("Successfully load config/.env")
}
