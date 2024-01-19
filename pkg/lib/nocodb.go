package lib

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/esc-chula/esc-docsync/platform/nocodb"
)

func FetchNocoDBTableData(tableName string) ([]map[string]interface{}, error) {
	log := logger.GetLogger()

	schema := config.GetDataSchema(tableName)

	nocodbService := nocodb.NewNocoDBService()

	tableId := config.GetNocoDBTableId(tableName)

	rows, err := nocodbService.GetRows(tableId)
	if err != nil {
		log.Error("Failed to get table " + tableId)
		return nil, err
	}

	mappedResults := make([]map[string]interface{}, 0)

	for _, row := range rows {
		mappedProperties := make(map[string]interface{})
		mappedProperties["Id"] = row["Id"]
		mappedProperties["Notion Page Id"] = row["Notion Page Id"]

		for _, schema := range schema {
			if schema.NotionType == "" || schema.NocoDBType == "" {
				continue
			}

			if row[schema.Name] == nil {
				continue
			}

			mappedProperties[schema.Name] = row[schema.Name]
		}

		mappedResults = append(mappedResults, mappedProperties)
	}

	return mappedResults, nil
}

func UpdateNocoDBRows(data []map[string]interface{}, tableName string) error {
	log := logger.GetLogger()

	nocodbService := nocodb.NewNocoDBService()

	tableId := config.GetNocoDBTableId(tableName)

	if err := nocodbService.UpdateRows(tableId, data); err != nil {
		log.Error("Failed to update rows for table " + tableId)
		return err
	}

	return nil
}
