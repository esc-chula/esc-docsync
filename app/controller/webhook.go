package controller

import (
	"github.com/esc-chula/esc-docsync/app/model"
	"github.com/esc-chula/esc-docsync/pkg/data"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/esc-chula/esc-docsync/platform/nocodb"
	"github.com/esc-chula/esc-docsync/platform/notion"
	"github.com/esc-chula/esc-docsync/platform/notion/template"
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
			return c.JSON(fiber.Map{
				"message": "Database ID not found for table: " + body.Data.TableName,
			})
		}

		pageData, err := notionService.CreatePage(
			template.ReqPageBody(databaseId, fiber.Map{
				"Title":       template.TitleProperty(row.Title),
				"Description": template.RichTextProperty(row.Description),
			}),
		)
		if err != nil {
			log.Error(err)
			return err
		}

		log.Info("Successfully created Notion page: ", pageData.Id)

		nocodbService := nocodb.NewNocoDBService()
		tableId := data.GetNocoDBTableId(body.Data.TableName)
		if tableId == "" {
			log.Error("Table ID not found for table: ", body.Data.TableName)
			return c.JSON(fiber.Map{
				"message": "Table ID not found for table: " + body.Data.TableName,
			})
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

	return c.JSON(fiber.Map{
		"message": "Webhook received",
	})
}
