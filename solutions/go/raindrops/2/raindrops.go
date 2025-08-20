package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var b strings.Builder

	if number%3 == 0 {
		b.WriteString("Pling")
	}

	if number%5 == 0 {
		b.WriteString("Plang")
	}

	if number%7 == 0 {
		b.WriteString("Plong")
	}

	if r := b.String(); r != "" {
		return r
	}

	return strconv.Itoa(number)
}
