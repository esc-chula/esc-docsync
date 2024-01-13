package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           time.Second * time.Duration(AppConfig().ReadTimeout),
	}
}
