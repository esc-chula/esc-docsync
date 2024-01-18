package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/esc-chula/esc-docsync/platform/logger"
)

type Schema struct {
	Name       string `json:"name"`
	NocoDBType string `json:"nocodb_type"`
	NotionType string `json:"notion_type"`
}

type DataMapModel struct {
	TableName        string   `json:"table_name"`
	NocoDBTableId    string   `json:"nocodb_table_id"`
	NotionDatabaseId string   `json:"notion_database_id"`
	Schema           []Schema `json:"schema"`
}

func ReadDataMap() ([]byte, error) {
	log := logger.GetLogger()

	jsonFile, err := os.Open("config/data_map.json")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	return byteValue, nil
}

func GetDataMap() []DataMapModel {
	byteValue, err := ReadDataMap()
	if err != nil {
		return nil
	}

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

func GetSchema(tableName string) []Schema {
	dataMap := GetDataMap()

	for _, data := range dataMap {
		if data.TableName == tableName {
			return data.Schema
		}
	}

	return nil
}
