package main

import (
	"github.com/esc-chula/esc-docsync/cmd/job"
	"github.com/esc-chula/esc-docsync/cmd/webhook"
	"github.com/esc-chula/esc-docsync/pkg/config"
)

func main() {
	config.LoadAllConfigs()

	job.RegisterJobs()

	webhook.Serve()
}
