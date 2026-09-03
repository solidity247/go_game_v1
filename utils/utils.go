package utils

import (
	"slices"
	"strings"
)

func JoinValues(delimiter string, values ...string) string {
	vals := slices.Collect(func(yield func(string) bool) {
		for _, v := range values {
			if v == "" {
				continue
			}
			if !yield(v) {
				return
			}
		}
	})
	return strings.Join(vals, delimiter)
}
