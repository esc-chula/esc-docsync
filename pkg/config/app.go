package config

import (
	"os"
)

type App struct {
	NotionAPIKey     string
	NotionDatabaseID string
	NocoDBAPIToken   string
}

var app = &App{}

func AppConfig() *App {
	return app
}

func LoadAppConfig() {
	app.NotionAPIKey = os.Getenv("NOTION_API_KEY")
	app.NotionDatabaseID = os.Getenv("NOTION_DATABASE_ID")
	app.NocoDBAPIToken = os.Getenv("NOCODB_API_TOKEN")
}
