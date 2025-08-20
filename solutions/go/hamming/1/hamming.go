package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Cannot calculate Hamming distance between samples or non-equal length.")
	}

	var count int

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
