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

		if isDataValid := lib.ValidateDataChanges(nocodbData, notionData); isDataValid {
			continue
		}

		if len(nocodbData) != len(notionData) {
			log.Info("Data length is not equal")
			continue
		} else {
			unsyncedNotionPageIds := lib.FindUnsyncedNocoDBNotionPageIds(nocodbData, notionData)

			updatingData := make([]map[string]interface{}, 0)

			for _, notionRow := range notionData {
				for _, unsyncedNotionPageId := range unsyncedNotionPageIds {
					var nocodbRowId float64

					for _, row := range nocodbData {
						if row["Notion Page Id"] == unsyncedNotionPageId {
							nocodbRowId = row["Id"].(float64)
						}
					}

					if notionRow["Notion Page Id"] == unsyncedNotionPageId {
						delete(notionRow, "Notion Page Id")
						notionRow["Id"] = nocodbRowId
						updatingData = append(updatingData, notionRow)
					}
				}
			}

			if err := lib.UpdateNocoDBRows(updatingData, data.TableName); err != nil {
				log.Error(err)
				continue
			}
		}

		log.Info("Data successfully synced")
	}
}
