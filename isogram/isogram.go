package isogram

// IsIsogram check is a word is an isogram, meaning has no repeatable words.
//
// Modified to combine fant0mz's optimized version.
//
// Original version: https://exercism.org/tracks/go/exercises/isogram/solutions/thinkverse
//
// Optimized version: https://exercism.org/tracks/go/exercises/isogram/solutions/fant0mz
//
// Slightly more optimized version using a bit mask, thank you for the inspiration @bobahop
//
// https://exercism.org/profiles/bobahop/
func IsIsogram(word string) bool {
	var mask uint32

	for i := 0; i < len(word); i++ {
		char := byte(word[i])

		// Deal only with lowercase letters.
		if char <= 'Z' {
			char += 1 << 5
		}

		// Make sure we only add lowercase letters.
		if 'a' <= char && char <= 'z' {
			// Check if the character is recorded.
			if (mask | (1 << (char - 97))) == mask {
				return false
			}

			// Record the character in the bit mask.
			mask |= 1 << (char - 97)
		}
	}

	return true
}
