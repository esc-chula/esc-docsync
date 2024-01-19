package nocodb

import (
	"encoding/json"

	"github.com/esc-chula/esc-docsync/platform/nocodb/model"
	"github.com/gofiber/fiber/v2"
)

func (r *NocoDBService) GetRows(tableId string) ([]map[string]interface{}, error) {
	client := NocoDBHTTPClient()

	resp, err := client.Get("/tables/" + tableId + "/records")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res model.ResRowsBody
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	rows := res.List

	return rows, nil
}

func (r *NocoDBService) GetRow() {
}

func (r *NocoDBService) CreateRow() {
}

func (r *NocoDBService) UpdateRow(tableId string, data fiber.Map) error {
	client := NocoDBHTTPClient()

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := client.Patch("/tables/"+tableId+"/records", "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
