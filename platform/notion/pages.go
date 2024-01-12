package notion

import (
	"encoding/json"

	"github.com/esc-chula/esc-docsync/platform/notion/template"
	"github.com/gofiber/fiber/v2"
)

func (r *NotionService) CreatePage(data fiber.Map) (template.ResPageBody, error) {
	client := NotionHTTPClient()

	body, err := json.Marshal(data)
	if err != nil {
		return template.ResPageBody{}, err
	}

	resp, err := client.Post("/pages", "application/json", body)
	if err != nil {
		return template.ResPageBody{}, err
	}
	defer resp.Body.Close()

	var res template.ResPageBody
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return template.ResPageBody{}, err
	}

	return res, nil
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
