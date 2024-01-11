package data

import (
	"encoding/json"
	"io"
	"os"

	"github.com/esc-chula/esc-docsync/platform/logger"
)

type DataMapModel struct {
	TableName        string `json:"table_name"`
	NocoDBTableId    string `json:"nocodb_table_id"`
	NotionDatabaseId string `json:"notion_database_id"`
}

func GetDataMap() []DataMapModel {
	log := logger.GetLogger()

	jsonFile, err := os.Open("config/data_map.json")
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var dataMap []DataMapModel

	json.Unmarshal(byteValue, &dataMap)

	return dataMap
}

func GetNocoDBTableId(tableName string) string {
	dataMap := GetDataMap()

	for _, data := range dataMap {
		if data.TableName == tableName {
			return data.NocoDBTableId
		}
	}

	return ""
}

func GetNotionDatabaseId(tableName string) string {
	dataMap := GetDataMap()

	for _, data := range dataMap {
		if data.TableName == tableName {
			return data.NotionDatabaseId
		}
	}

	return ""
}
