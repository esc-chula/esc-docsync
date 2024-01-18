package route

import (
	"github.com/esc-chula/esc-docsync/app/controller"
	"github.com/esc-chula/esc-docsync/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func WebhookRoutes(app *fiber.App) {
	webhookRoutes := app.Group("/webhook", middleware.WebhookGuard)

	webhookRoutes.Post("/insert", controller.InsertHandler)
	webhookRoutes.Post("/update", controller.UpdateHandler)
}
