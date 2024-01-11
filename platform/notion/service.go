package notion

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/http"
)

type NotionService struct{}

func NewNotionService() *NotionService {
	return &NotionService{}
}

func NotionHTTPClient() *http.CustomHTTPClient {
	notionConfig := config.NotionConfig()

	client := http.NewCustomHTTPClient()
	client.SetBaseURL("https://api.notion.com/v1")
	client.SetDefaultHeaders(map[string]string{
		"Content-Type":   "application/json",
		"Notion-Version": "2022-06-28",
	})
	client.SetBearerAuthorization(notionConfig.NotionAPIKey)

	return client
}
