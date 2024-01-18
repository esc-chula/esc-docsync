package webhook

import (
	"github.com/esc-chula/esc-docsync/app/route"
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New(config.FiberConfig())

	middleware.FiberMiddleware(app)

	route.GeneralRoute(app)
	route.WebhookRoutes(app)
	route.NotFoundRoute(app)

	SeverHandler(app)
}
