package code

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return 0
	}

	if len(needle) == 0 {
		return  0
	}

	// return  strings.Index(haystack, needle)

	for i := 0; i <= len(haystack) - len(needle); i++ {
		j := 0

		for  j < len(haystack) && haystack[i+j] == needle[j] {
			j++
		}

		if j == len(needle) {
			return  i
		}
	}


    return -1
}