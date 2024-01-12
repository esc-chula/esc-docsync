package nocodb

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (r *NocoDBService) GetRows() {
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
