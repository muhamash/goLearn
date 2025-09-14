package code

import "strings"

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return 0
	}

	if len(needle) == 0 {
		return  0
	}

	return  strings.Index(haystack, needle)
}