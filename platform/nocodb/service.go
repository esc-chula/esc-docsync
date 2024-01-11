package nocodb

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/http"
)

type NocoDBService struct {
}

func NewNocoDBService() *NocoDBService {
	return &NocoDBService{}
}

func NocoDBHTTPClient() *http.CustomHTTPClient {
	appConfig := config.AppConfig()

	client := http.NewCustomHTTPClient()
	client.SetBaseURL("http://localhost:8080/api/v2")
	client.SetDefaultHeaders(map[string]string{
		"Content-Type":   "application/json",
		"NocoDB-Version": "2022-06-28",
	})
	client.SetBearerAuthorization(appConfig.NocoDBAPIToken)

	return client
}
