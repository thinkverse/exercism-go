package pangram

// Optimized algorithm using bit fields.
// Thank you for the inspiration @bobahop
//
// https://exercism.org/profiles/bobahop/
func IsPangram(input string) bool {
	var mask uint32

	for i := 0; i < len(input); i++ {
		char := byte(input[i])

		if char <= 'Z' {
			char += 1 << 5
		}

		if 'a' <= char && char <= 'z' {
			mask |= 1 << (char - 97)
		}
	}

	return mask == ((1 << 26) - 1)
}
