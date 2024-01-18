package job

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/lib"
	"github.com/esc-chula/esc-docsync/platform/logger"
)

func FetchNotionCalendar() {
}

func FetchNotionData() {
	log := logger.GetLogger()
	dataMap := config.GetDataMap()

	for _, data := range dataMap {
		data, err := lib.FetchNotionDatabase(data.TableName)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Info(data[0]["Title"])
	}
}
