package isogram

func IsIsogram(word string) bool {
	chars := map[rune]int{}

	for _, char := range word {
		// Deal only with uppercase letters.
		if 97 <= char && char <= 122 {
			char -= 32
		}

		// Make sure we only add uppercase letters.
		if 65 <= char && char <= 90 {
			chars[char]++
		}

		// Check if we've seen the char before.
		if seen := chars[char]; seen > 1 {
			return false
		}
	}

	return true
}
