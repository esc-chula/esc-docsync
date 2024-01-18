package webhook

import (
	"github.com/esc-chula/esc-docsync/app/routes"
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New(config.FiberConfig())

	middleware.FiberMiddleware(app)

	routes.GeneralRoute(app)
	routes.WebhookRoutes(app)
	routes.NotFoundRoute(app)

	SeverHandler(app)
}
