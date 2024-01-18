package config

import (
	"github.com/esc-chula/esc-docsync/pkg/calendar"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/joho/godotenv"
)

func LoadENV(envFile string) error {
	err := godotenv.Load(envFile)
	if err != nil {
		return err
	}

	LoadAppConfig()
	LoadNocoDBConfig()
	LoadNotionConfig()

	return nil
}

func LoadAllConfigs() {
	log := logger.GetLogger()

	var err error

	err = LoadENV("config/.env")
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}
	log.Info("Successfully load config/.env")

	_, err = ReadMap("config/data_map.json")
	if err != nil {
		log.Fatalf("can't load config/data_map.json. error: %v", err)
	}
	log.Info("Successfully load config/data_map.json")

	_, err = ReadMap("config/calendar_map.json")
	if err != nil {
		log.Fatalf("can't load config/calendar_map.json. error: %v", err)
	}
	log.Info("Successfully load config/calendar_map.json")

	_, err = calendar.ReadSecret()
	if err != nil {
		log.Fatalf("can't load config/google_secret.json. error: %v", err)
	}
	log.Info("Successfully load config/google_secret.json")

	_, err = calendar.ReadToken()
	if err != nil {
		log.Fatalf("can't load config/google_token.json. error: %v", err)
	}
	log.Info("Successfully load config/google_token.json")
}
