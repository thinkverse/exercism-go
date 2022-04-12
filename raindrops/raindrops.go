package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	var b string

	if number%3 == 0 {
		b += "Pling"
	}

	if number%5 == 0 {
		b += "Plang"
	}

	if number%7 == 0 {
		b += "Plong"
	}

	if b != "" {
		return b
	}

	return strconv.Itoa(number)
}
