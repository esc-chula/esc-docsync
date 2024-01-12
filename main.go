package main

import (
	"github.com/esc-chula/esc-docsync/cmd/data"
	"github.com/esc-chula/esc-docsync/cmd/webhook"
	"github.com/esc-chula/esc-docsync/pkg/config"
)

func main() {
	config.LoadAllConfigs("config/.env")

	data.LoadJsonConfigs()

	webhook.Serve()
}
