package controller

import (
	"github.com/esc-chula/esc-docsync/app/model"
	"github.com/esc-chula/esc-docsync/pkg/data"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/esc-chula/esc-docsync/platform/nocodb"
	"github.com/esc-chula/esc-docsync/platform/notion"
	"github.com/gofiber/fiber/v2"
)

func InsertHandler(c *fiber.Ctx) error {
	log := logger.GetLogger()

	var body model.InsertBody
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return err
	}

	notionService := notion.NewNotionService()

	for _, row := range body.Data.Rows {
		databaseId := data.GetNotionDatabaseId(body.Data.TableName)

		if databaseId == "" {
			log.Error("Database ID not found for table: ", body.Data.TableName)
			return c.SendString("Database ID not found for table: " + body.Data.TableName)
		}

		pageData, err := notionService.CreatePage(fiber.Map{
			"parent": fiber.Map{
				"database_id": databaseId,
			},
			"properties": fiber.Map{
				"Title": fiber.Map{
					"title": []fiber.Map{
						{
							"text": fiber.Map{
								"content": row.Title,
							},
						},
					},
				},
				"Description": fiber.Map{
					"rich_text": []fiber.Map{
						{
							"text": fiber.Map{
								"content": row.Description,
							},
						},
					},
				},
			},
		})
		if err != nil {
			log.Error(err)
			return err
		}

		log.Info("Successfully created Notion page: ", pageData.Id)

		nocodbService := nocodb.NewNocoDBService()
		tableId := data.GetNocoDBTableId(body.Data.TableName)
		if tableId == "" {
			log.Error("Table ID not found for table: ", body.Data.TableName)
			return c.SendString("Table ID not found for table: " + body.Data.TableName)
		}

		if err := nocodbService.UpdateRow(tableId, fiber.Map{
			"Id":             row.Id,
			"Notion Page Id": pageData.Id,
		}); err != nil {
			log.Error(err)
			return err
		}

		log.Info("Successfully updated NocoDB record: ", row.Id)
	}

	return c.SendString("Hello, World 👋!")
}
