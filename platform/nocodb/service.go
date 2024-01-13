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
	nocodbConfig := config.NocoDBConfig()

	client := http.NewCustomHTTPClient()
	client.SetBaseURL(nocodbConfig.NocoDBURL)
	client.SetDefaultHeaders(map[string]string{
		"Content-Type": "application/json",
		"xc-token":     nocodbConfig.NocoDBAPIToken,
	})
	client.SetBearerAuthorization(nocodbConfig.NocoDBAPIToken)

	return client
}
