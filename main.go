package main

import (
	"github.com/esc-chula/esc-docsync/cmd/webhook"
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/pkg/data"
)

func main() {
	config.LoadAllConfigs("config/.env")

	data.GetDataMap()

	webhook.Serve()
}
