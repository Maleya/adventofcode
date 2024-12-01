// Package stringutil provides extended functionality for []string.
package stringsliceutil

func contains(search string, slice []string) bool {
	for _, element := range slice {
		if element == search {
			return true
		}
	}
	return false
}
