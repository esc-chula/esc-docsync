package lib

import (
	"fmt"

	"github.com/esc-chula/esc-docsync/pkg/util"
)

func ValidateCacheChanges(data []map[string]interface{}, key string) bool {
	return false
}

func ValidateDataChanges(dataA, dataB []map[string]interface{}) bool {
	sortedDataA := util.SortMapInterface(dataA, "Title")
	sortedDataB := util.SortMapInterface(dataB, "Title")

	zippedData, err := util.ZipMapInterface(sortedDataA, sortedDataB)
	if err != nil {
		return false
	}

	var equal bool

	for _, data := range zippedData {
		dataAKeys := util.Keys(data.A)
		dataBKeys := util.Keys(data.B)

		keys := util.FindMatchedKeys(dataAKeys, dataBKeys)

		for _, key := range keys {
			if data.A[key] == data.B[key] {
				equal = true
			} else {
				equal = false
			}
		}
	}

	return equal
}

func FindUnsyncedNocoDBNotionPageIds(dataA, dataB []map[string]interface{}) []string {
	sortedDataA := util.SortMapInterface(dataA, "Title")
	sortedDataB := util.SortMapInterface(dataB, "Title")

	zippedData, err := util.ZipMapInterface(sortedDataA, sortedDataB)
	if err != nil {
		return nil
	}

	var unsyncedNocoDBNotionPageIds []string

	for _, data := range zippedData {
		dataAKeys := util.Keys(data.A)
		dataBKeys := util.Keys(data.B)

		keys := util.FindMatchedKeys(dataAKeys, dataBKeys)

		for _, key := range keys {
			if data.A[key] != data.B[key] {
				unsyncedNocoDBNotionPageIds = append(unsyncedNocoDBNotionPageIds, fmt.Sprintf("%v", data.A["Notion Page Id"]))
			}
		}
	}

	return unsyncedNocoDBNotionPageIds
}
