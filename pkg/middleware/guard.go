package middleware

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func WebhookGuard(c *fiber.Ctx) error {
	appConfig := config.AppConfig()

	if c.Get("X-ESC-Secret") != appConfig.SecretKey {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
