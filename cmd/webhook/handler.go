package webhook

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/platform/logger"
	"github.com/gofiber/fiber/v2"
)

func SeverHandler(app *fiber.App) {
	appConfig := config.AppConfig()
	log := logger.GetLogger()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		log.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	serverAddr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Errorf("Oops... server is not running! error: %v", err)
	}
}
