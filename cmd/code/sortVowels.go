package code

import (
	"sort"
	"strings"
)

func sortVowels(s string) string {
	isVowel := func(c byte) bool {
		return strings.ContainsRune("aeiouAEIOU", rune(c))
	}

	vowels := []byte{}
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]){
			vowels = append(vowels,s[i])
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return  vowels[i] < vowels[j]
	})

	result := make([]byte, len(s))
	vowelsIdx := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]){
			result[i] = vowels[vowelsIdx]
			vowelsIdx++
		} else {
			result[i] = s[i]
		}
	}

	return string(result)
}