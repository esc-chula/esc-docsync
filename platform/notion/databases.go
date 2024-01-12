package notion

import (
	"fmt"
)

func (r *NotionService) RetrieveDatabase(databaseId string) {
	client := NotionHTTPClient()

	resp, err := client.Get("/databases/" + databaseId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}

func (r *NotionService) QueryDatabase(databaseId string) {

	client := NotionHTTPClient()

	resp, err := client.Post("/databases/"+databaseId+"/query", "application/json", []byte(`{}`))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
