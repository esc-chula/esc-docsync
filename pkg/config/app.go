package config

import (
	"os"
)

type App struct {
	NotionAPIKey     string
	NotionDatabaesID string
	NocoDBAPIToken   string
}

var app = &App{}

func AppCfg() *App {
	return app
}

func LoadAppConfig() {
	app.NotionAPIKey = os.Getenv("NOTION_API_KEY")
	app.NotionDatabaesID = os.Getenv("NOTION_DATABASE_ID")
	app.NocoDBAPIToken = os.Getenv("NOCODB_API_TOKEN")
}
