package route

import (
	"github.com/esc-chula/esc-docsync/app/controller"
	"github.com/gofiber/fiber/v2"
)

func WebhookRoutes(app *fiber.App) {
	webhookRoutes := app.Group("/webhook")

	webhookRoutes.Post("/insert", controller.InsertHandler)
	webhookRoutes.Post("/update", controller.UpdateHandler)
}
