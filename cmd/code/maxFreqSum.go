package code


func maxFreqSum(s string) int {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	freq := make(map[rune] int)

	for _, ch := range s {
		freq[ch]++
	}

	maxVowels, maxConsonant := 0,0

	for ch, count := range freq {
		if vowels[ch] {
			if count > maxVowels {
				maxVowels = count
			}
		} else {
			if count > maxConsonant {
				maxConsonant = count
			}
		}
	}

 
	return  maxVowels + maxConsonant
}