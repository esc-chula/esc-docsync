package data

import (
	"github.com/esc-chula/esc-docsync/pkg/data"
	"github.com/esc-chula/esc-docsync/platform/logger"
)

func LoadJsonConfigs() {
	log := logger.GetLogger()

	data.ReadDataMap()
	log.Info("Successfully read config/data/data_map.json")
}
