package util

import "sort"

func SortString(data []string) []string {
	compareTitles := func(i, j int) bool {
		titleI := data[i]
		titleJ := data[j]
		return titleI < titleJ
	}

	sort.Slice(data, compareTitles)

	return data
}

func SortMapInterface(data []map[string]interface{}, sortBy string) []map[string]interface{} {
	compareTitles := func(i, j int) bool {
		titleI := data[i][sortBy].(string)
		titleJ := data[j][sortBy].(string)
		return titleI < titleJ
	}

	sort.Slice(data, compareTitles)

	return data
}
