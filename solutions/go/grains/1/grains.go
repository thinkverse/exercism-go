package grains

import (
	"errors"
)

// 8500328 ~139 ns/op
func Square(number int) (uint64, error) {
	if 0 >= number || number >= 65 {
		return 0, errors.New("Number must be in a range of 1 to 64.")
	}

	if number == 1 {
		return 1, nil
	}

	var r uint64

	for i := 0; i <= (number - 2); i++ {
		r = 2 << i
	}

	return r, nil
}

// Far shy from being optimized for
// speed, but for a first iteration
//
// 354817 ~3360 ns/op
func Total() uint64 {
	var t uint64

	for i := 1; i <= 64; i++ {
		n, _ := Square(i)
		t += n
	}

	return t
}
