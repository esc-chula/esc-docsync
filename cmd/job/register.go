package job

import (
	"github.com/esc-chula/esc-docsync/app/job"
	"github.com/robfig/cron"
)

func RegisterJobs() {
	c := cron.New()

	c.AddFunc("@every 30s", job.SyncNotionToNocoDB)

	c.Start()
}
