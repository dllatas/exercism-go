package etl

import (
	"strings"
)

func Transform(a given) expectation {
	var result map[string]int
	result = make(map[string]int)
	for k, v := range a {
		for _, w := range v {
			result[strings.ToLower(w)] = k
		}
	}
	return result
}
