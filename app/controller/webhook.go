package controller

import (
	"github.com/esc-chula/esc-docsync/app/model"
	"github.com/esc-chula/esc-docsync/pkg/data"
	"github.com/esc-chula/esc-docsync/platform/logger"
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

	log.Info(body.Data.TableName)
	log.Info(data.GetNotionDatabaseId(body.Data.TableName))

	notionService := notion.NewNotionService()

	for _, row := range body.Data.Rows {
		err := notionService.CreatePage(fiber.Map{
			"parent": fiber.Map{
				"database_id": data.GetNotionDatabaseId(body.Data.TableName),
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
	}

	log.Info("Successfully created Notion page.")

	return c.SendString("Hello, World 👋!")
}
