package config

import "encoding/json"

type CalendarMapModel struct {
	CalendarId       string `json:"calendar_id"`
	NotionDatabaseId string `json:"notion_database_id"`
}

func GetCalendarMap() []CalendarMapModel {
	byteValue, err := ReadMap("config/calendar_map.json")
	if err != nil {
		return nil
	}

	var calendarMap []CalendarMapModel

	json.Unmarshal(byteValue, &calendarMap)

	return calendarMap
}
