package main

import (
	"github.com/esc-chula/esc-docsync/cmd/webhook"
	"github.com/esc-chula/esc-docsync/pkg/config"
)

func main() {
	config.LoadAllConfigs()

	webhook.Serve()
}
