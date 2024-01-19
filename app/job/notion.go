package job

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/lib"
	"github.com/esc-chula/esc-docsync/platform/logger"
)

func SyncNotionToNocoDB() {
	log := logger.GetLogger()
	dataMap := config.GetDataMap()

	for _, data := range dataMap {
		notionData, err := lib.FetchNotionDatabaseData(data.TableName)
		if err != nil {
			log.Error(err)
			continue
		}

		isCacheValid := lib.ValidateCacheChanges(notionData, data.TableName)

		if isCacheValid {
			log.Info("Cache is valid")
			continue
		}

		nocodbData, err := lib.FetchNocoDBTableData(data.TableName)
		if err != nil {
			log.Error(err)
			continue
		}

		isDataValid, err := lib.ValidateDataChanges(nocodbData, notionData)
		if err != nil {
			log.Error(err)
			continue
		}

		if isDataValid {
			log.Info("Data already synced")
			continue
		}

		if len(nocodbData) != len(notionData) {
			log.Info("Data length is not equal")
			continue
		}

		// err = lib.SyncData(notionData, data.TableName)
		// if err != nil {
		// 	log.Error(err)
		// 	continue
		// }

		log.Info("Data successfully synced")
	}
}
