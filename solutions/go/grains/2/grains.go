package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if 0 >= number || number >= 65 {
		return 0, errors.New("Number must be in a range of 1 to 64.")
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
