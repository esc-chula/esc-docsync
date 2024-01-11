package notion

import (
	"fmt"

	"github.com/esc-chula/esc-docsync/pkg/config"
)

func (r *NotionService) RetrieveDatabase() {
	notionConfig := config.NotionConfig()

	client := NotionHTTPClient()

	resp, err := client.Get("/databases/" + notionConfig.NotionDatabaseID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}

func (r *NotionService) QueryDatabase() {
	notionConfig := config.NotionConfig()

	client := NotionHTTPClient()

	resp, err := client.Post("/databases/"+notionConfig.NotionDatabaseID+"/query", "application/json", []byte(`{}`))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
