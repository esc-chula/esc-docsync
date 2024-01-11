package notion

import (
	"fmt"
)

func (r *NotionService) CreatePage(data []byte) {
	client := NotionHTTPClient()

	resp, err := client.Post("/pages", "application/json", data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}

func (r *NotionService) RetrievePage(pageId string) {
	client := NotionHTTPClient()

	resp, err := client.Get("/pages/" + pageId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}

func (r *NotionService) UpdatePage(pageId string, data []byte) {
	client := NotionHTTPClient()

	resp, err := client.Patch("/pages/"+pageId, "application/json", data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
