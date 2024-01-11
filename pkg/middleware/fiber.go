package middleware

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var log = logger.GetLogger()

func FiberMiddleware(app *fiber.App) {
	appConfig := config.AppConfig()

	app.Use(
		cors.New(cors.Config{
			AllowOrigins: appConfig.AllowedOrigin,
			AllowMethods: "POST",
		}),

		RequestLogger(log),
	)
}
