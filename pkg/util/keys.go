package util

func Keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func FindMatchedKeys(a, b []string) []string {
	var keys []string

	for _, keyA := range a {
		for _, keyB := range b {
			if keyA == keyB {
				keys = append(keys, keyA)
			}
		}
	}

	sortedKeys := SortString(keys)

	return sortedKeys
}
