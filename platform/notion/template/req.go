package template

import "github.com/gofiber/fiber/v2"

func ReqPageBody(databaseId string, properties fiber.Map) fiber.Map {
	return fiber.Map{
		"parent": fiber.Map{
			"database_id": databaseId,
		},
		"properties": properties,
	}
}

func TitleProperty(title string) fiber.Map {
	return fiber.Map{
		"title": []fiber.Map{
			{
				"text": fiber.Map{
					"content": title,
				},
			},
		},
	}
}

func RichTextProperty(text string) fiber.Map {
	return fiber.Map{
		"rich_text": []fiber.Map{
			{
				"text": fiber.Map{
					"content": text,
				},
			},
		},
	}
}
