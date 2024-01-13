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
	client.SetBaseURL(notionConfig.NotionURL)
	client.SetDefaultHeaders(map[string]string{
		"Content-Type":   "application/json",
		"Notion-Version": notionConfig.NotionVersion,
	})
	client.SetBearerAuthorization(notionConfig.NotionAPIKey)

	return client
}
