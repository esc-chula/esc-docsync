package lib

import (
	"fmt"

	"github.com/esc-chula/esc-docsync/pkg/util"
)

func ValidateCacheChanges(data []map[string]interface{}, key string) bool {
	return false
}

func ValidateDataChanges(dataA, dataB []map[string]interface{}) (bool, error) {
	sortedDataA := util.SortMapInterface(dataA, "Title")
	sortedDataB := util.SortMapInterface(dataB, "Title")

	zippedData, err := util.ZipMapInterface(sortedDataA, sortedDataB)
	if err != nil {
		return false, nil
	}

	var equal bool
	var unequalCount int

	for _, data := range zippedData {
		dataAKeys := util.Keys(data.A)
		dataBKeys := util.Keys(data.B)

		keys := util.FindMatchedKeys(dataAKeys, dataBKeys)

		for _, key := range keys {
			if data.A[key] == data.B[key] {
				equal = true
			} else {
				equal = false
				unequalCount++
			}
		}
	}

	if equal {
		return true, nil
	}

	return false, fmt.Errorf("data is not equal, %d unequal property", unequalCount)
}
