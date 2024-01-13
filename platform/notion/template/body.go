package template

import "github.com/gofiber/fiber/v2"

func CreatePageBody(databaseId string, properties fiber.Map) fiber.Map {
	return fiber.Map{
		"parent": fiber.Map{
			"database_id": databaseId,
		},
		"properties": properties,
	}
}

func NotionTypeProperty(propertyType string, value string) fiber.Map {
	switch propertyType {
	case "title":
		return TitleProperty(value)
	case "rich_text":
		return RichTextProperty(value)
	}
	return RichTextProperty(value)
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
