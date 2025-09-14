package code

import "strings"

func repeatedSubstringPattern(s string) bool {
	doubleString := s + s

	return  strings.Contains(doubleString[1:len(doubleString)-1],s)
}