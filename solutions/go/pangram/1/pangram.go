package pangram

func IsPangram(input string) bool {
	set := make(map[rune]bool, 26)

	for _, char := range input {
		if char <= 'Z' {
			char += 1 << 5
		}

		if 'a' <= char && char <= 'z' {
			set[char] = true
		}
	}

	return len(set) == 26
}
