package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	var b strings.Builder

	if number%105 == 0 {
		b.WriteString("PlingPlangPlong")
	} else if number%35 == 0 {
		b.WriteString("PlangPlong")
	} else if number%21 == 0 {
		b.WriteString("PlingPlong")
	} else if number%15 == 0 {
		b.WriteString("PlingPlang")
	} else if number%3 == 0 {
		b.WriteString("Pling")
	} else if number%5 == 0 {
		b.WriteString("Plang")
	} else if number%7 == 0 {
		b.WriteString("Plong")
	} else {
		b.WriteString(fmt.Sprintf("%d", number))
	}

	return b.String()
}
