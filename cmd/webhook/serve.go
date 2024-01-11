package webhook

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/middleware"
	"github.com/esc-chula/esc-docsync/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New(config.FiberConfig())

	middleware.FiberMiddleware(app)

	route.WebhookRoutes(app)

	SeverHandler(app)
}
