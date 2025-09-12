package code

func DoesAliceWin(s string) bool {
    vowels := "aeiou"
	count := 0

	for _, ch := range s {
		if contains(vowels, ch) {
			count++
		}
	}

	return count > 0
}

func contains(vowels string, ch rune) bool {
	for _, v := range vowels {
		if v == ch {
			return true
		}
	}
	return false
}