package code

import "strings"

// Helper: check if rune is a vowel
func isVowel (ch rune) bool {
	return  ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

// Transform a word into its vowel-error key
func devowel(word string) string {
	var sb strings.Builder

	for _, ch:= range word {
		if isVowel(ch){
			sb.WriteRune('*')
		} else {
			sb.WriteRune(ch)
		}
	}

	return  sb.String()
}

// final function

func SpellChecker(wordlist []string, queries []string) []string {
	// exact match set
	exactWords := make(map[string] bool)

	// case insensitive match
	caseInsensitive := make(map[string] string)

	// vowel error map
	vowelInsensitive := make(map[string] string)

	for _, word := range wordlist {
		exactWords[word] = true

		lower := strings.ToLower(word)
		if _, exists := caseInsensitive[lower]; !exists {
			caseInsensitive[lower] = word
		} 

		dev := devowel(lower)

		if _, exists := vowelInsensitive[dev]; !exists {
			vowelInsensitive[dev] = word
		}
	}

	result := make([]string, len(queries))
	for i, q := range queries {
		// exact match
		if exactWords[q] {
			result[i] = q
			continue
		}

		lower := strings.ToLower(q)

		// case insensitive
		if val, exists := caseInsensitive[lower]; exists {
			result[i] = val
			continue
		} 

		// vowel insensitive
		dev := devowel(lower)
		if val, exists := vowelInsensitive[dev]; exists {
			result[i] = val
			continue
		}

		// no match
		result[i] = ""
	}

	return result
}