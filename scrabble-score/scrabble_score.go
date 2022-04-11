package scrabble

var scores = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

func Score(word string) int {
	points := 0

	for _, char := range word {
		// If the char is lowercase, uppercase is by removing 32.
		// This works because there is 32 between a and A.
		// See: https://go.dev/play/p/QU8rPrW9JCz
		if char > 96 {
			char -= 32
		}

		if score, ok := scores[char]; ok {
			points += score
		}
	}

	return points
}
