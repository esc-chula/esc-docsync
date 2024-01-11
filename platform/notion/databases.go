package notion

import (
	"fmt"

	"github.com/esc-chula/esc-docsync/pkg/config"
)

func (r *NotionService) RetrieveDatabase() {
	appConfig := config.AppConfig()

	client := NotionHTTPClient()

	resp, err := client.Get("/databases/" + appConfig.NotionDatabaseID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}

func (r *NotionService) QueryDatabase() {
	appConfig := config.AppConfig()

	client := NotionHTTPClient()

	resp, err := client.Post("/databases/"+appConfig.NotionDatabaseID+"/query", "application/json", []byte(`{}`))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
