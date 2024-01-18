package lib

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/esc-chula/esc-docsync/platform/notion"
)

func FetchNotionDatabase(tableName string) ([]map[string]interface{}, error) {
	log := logger.GetLogger()

	notionService := notion.NewNotionService()

	schema := config.GetDataSchema(tableName)

	databaseId := config.GetNotionDatabaseId(tableName)

	dbData, err := notionService.QueryDatabase(databaseId)
	if err != nil {
		log.Error("Failed to query database " + databaseId)
		return nil, err
	}

	mappedResults := make([]map[string]interface{}, 0)

	for _, result := range dbData.Results {
		properties := result.Properties

		propertiesMap := properties.(map[string]interface{})

		mappedProperties := make(map[string]interface{})
		mappedProperties["Notion Page Id"] = result.Id

		for _, schema := range schema {
			if schema.NotionType == "" || schema.NocoDBType == "" {
				continue
			}

			if propertiesMap[schema.Name] == nil {
				continue
			}

			property := propertiesMap[schema.Name].(map[string]interface{})[schema.NotionType].([]interface{})[0]

			propertyValue, err := notionService.GetPropertyValue(property, schema.NotionType)
			if err != nil {
				log.Error("Failed to parse property " + schema.Name)
				continue
			}

			mappedProperties[schema.Name] = propertyValue
		}

		mappedResults = append(mappedResults, mappedProperties)
	}

	return mappedResults, nil
}
