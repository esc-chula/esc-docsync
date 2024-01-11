package notion

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (r *NotionService) CreatePage(body fiber.Map) error {
	client := NotionHTTPClient()

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := client.Post("/pages", "application/json", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (r *NotionService) RetrievePage(pageId string) error {
	client := NotionHTTPClient()

	resp, err := client.Get("/pages/" + pageId)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (r *NotionService) UpdatePage(pageId string, data []byte) error {
	client := NotionHTTPClient()

	resp, err := client.Patch("/pages/"+pageId, "application/json", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
