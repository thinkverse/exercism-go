package hamming

import (
	"errors"
)

// Iteration 1 with range:    14205139  ~82 ns/op
//
// Iteration 2 with no range: 25425594  ~46 ns/op
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Cannot calculate Hamming distance between samples or non-equal length.")
	}

	var count int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
