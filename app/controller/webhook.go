package controller

import (
	"github.com/esc-chula/esc-docsync/app/model"
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/esc-chula/esc-docsync/platform/nocodb"
	"github.com/esc-chula/esc-docsync/platform/notion"
	"github.com/esc-chula/esc-docsync/platform/notion/template"
	"github.com/gofiber/fiber/v2"
)

func InsertHandler(c *fiber.Ctx) error {
	log := logger.GetLogger()
	notionService := notion.NewNotionService()
	nocodbService := nocodb.NewNocoDBService()

	log.Info("Insert webhook received")

	var body model.InsertBody
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	tableName := body.Data.TableName
	if tableName == "" {
		err := "Table name not found"
		log.Error(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	rows := body.Data.Rows
	if len(rows) == 0 {
		err := "No rows found"
		log.Error(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	for _, row := range rows {
		rowData := row.(map[string]interface{})

		schema := config.GetDataSchema(tableName)

		properties := make(fiber.Map)
		for _, schema := range schema {
			if schema.NotionType == "" || schema.NocoDBType == "" {
				continue
			}
			properties[schema.Name] = template.NotionTypeProperty(
				schema.NotionType, rowData[schema.Name].(string),
			)
		}

		databaseId := config.GetNotionDatabaseId(tableName)

		if databaseId == "" {
			err := "Notion Database ID not found for table: " + tableName
			log.Error(err)
			return c.Status(500).JSON(fiber.Map{
				"message": err,
			})
		}

		pageData, err := notionService.CreatePage(
			template.CreatePageBody(databaseId, properties),
		)
		if err != nil {
			log.Error(err)
			return c.Status(500).JSON(fiber.Map{
				"message": err,
			})
		}

		log.Info("Successfully created Notion page: ", pageData.Id)

		tableId := config.GetNocoDBTableId(tableName)
		if tableId == "" {
			err := "NocoDB Table ID not found for table: " + tableName
			log.Error(err)
			return c.Status(500).JSON(fiber.Map{
				"message": err,
			})
		}

		if err := nocodbService.UpdateRow(tableId, fiber.Map{
			"Id":             rowData["Id"],
			"Notion Page Id": pageData.Id,
		}); err != nil {
			log.Error(err)
			return c.Status(500).JSON(fiber.Map{
				"message": err,
			})
		}

		log.Info("Successfully updated NocoDB record: ", rowData["Id"])
	}

	return c.JSON(fiber.Map{
		"message": "Webhook received",
	})
}

func UpdateHandler(c *fiber.Ctx) error {
	log := logger.GetLogger()

	log.Info("Update webhook received")

	var body model.UpdateBody
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
