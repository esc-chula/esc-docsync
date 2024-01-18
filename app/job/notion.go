package job

import (
	"github.com/esc-chula/esc-docsync/pkg/config"
	"github.com/esc-chula/esc-docsync/platform/logger"
)

func FetchNotionCalendar() {
	log := logger.GetLogger()

	calendarMap := config.GetCalendarMap()

	log.Info(calendarMap)
}

func FetchNotionData() {}
