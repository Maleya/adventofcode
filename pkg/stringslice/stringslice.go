// Package stringutil provides extended functionality for []string.
package stringslice

import (
	"log/slog"
	"strconv"
)

func Contains(search string, slice []string) bool {
	for _, element := range slice {
		if element == search {
			return true
		}
	}
	return false
}

func AtoiSlice(d []string) []int {
	intSlice := make([]int, len(d))
	for i, e := range d {
		convertedElement, err := strconv.Atoi(e)
		if err != nil {
			slog.Error("error converting string to int", "error", err)
		}
		intSlice[i] = convertedElement
	}
	return intSlice
}
