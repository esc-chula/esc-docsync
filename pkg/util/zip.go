package util

import "fmt"

type MapTuple struct {
	A, B map[string]interface{}
}

func ZipMapInterface(a, b []map[string]interface{}) ([]MapTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of the same length")
	}

	result := make([]MapTuple, len(a))

	for i := range a {
		result[i] = MapTuple{a[i], b[i]}
	}

	return result, nil
}
