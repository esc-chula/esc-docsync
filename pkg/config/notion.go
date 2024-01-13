package config

import (
	"os"
)

type Notion struct {
	NotionURL        string
	NotionVersion    string
	NotionAPIKey     string
	NotionDatabaseID string
}

var notion = &Notion{}

func NotionConfig() *Notion {
	return notion
}

func LoadNotionConfig() {
	notion.NotionURL = os.Getenv("NOTION_URL")
	notion.NotionVersion = os.Getenv("NOTION_VERSION")
	notion.NotionAPIKey = os.Getenv("NOTION_API_KEY")
	notion.NotionDatabaseID = os.Getenv("NOTION_DATABASE_ID")
}
