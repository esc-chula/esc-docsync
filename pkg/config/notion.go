package config

import (
	"os"
)

type Notion struct {
	NotionAPIKey     string
	NotionDatabaseID string
	NocoDBAPIToken   string
}

var notion = &Notion{}

func NotionConfig() *Notion {
	return notion
}

func LoadNotionConfig() {
	notion.NotionAPIKey = os.Getenv("NOTION_API_KEY")
	notion.NotionDatabaseID = os.Getenv("NOTION_DATABASE_ID")
}
