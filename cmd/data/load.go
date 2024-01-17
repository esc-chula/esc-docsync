package data

import (
	"github.com/esc-chula/esc-docsync/pkg/calendar"
	"github.com/esc-chula/esc-docsync/pkg/data"
	"github.com/esc-chula/esc-docsync/platform/logger"
)

func LoadJsonConfigs() {
	log := logger.GetLogger()

	data.ReadDataMap()
	log.Info("Successfully load config/data_map.json")

	calendar.ReadSecret()
	log.Info("Successfully load config/google_secret.json")

	calendar.ReadToken()
	log.Info("Successfully load config/google_token.json")
}
