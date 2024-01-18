package notion

import (
	"encoding/json"
	"fmt"

	"github.com/esc-chula/esc-docsync/platform/notion/model"
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

func (r *NotionService) QueryDatabase(databaseId string) (model.ResQueryDatabaseBody, error) {

	client := NotionHTTPClient()

	resp, err := client.Post("/databases/"+databaseId+"/query", "application/json", []byte(`{}`))
	if err != nil {
		return model.ResQueryDatabaseBody{}, err
	}
	defer resp.Body.Close()

	var res model.ResQueryDatabaseBody
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return model.ResQueryDatabaseBody{}, err
	}

	return res, nil
}
