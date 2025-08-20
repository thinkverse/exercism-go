package luhn

import (
	"regexp"
	"strconv"
)

func Valid(id string) bool {
	re := regexp.MustCompile(`\s+`)

	id = re.ReplaceAllString(id, "")

	length := len(id)

	if length <= 1 {
		return false
	}

	parity := (length - 2) % 2

	var product int

	for i, v := range id {
		num, err := strconv.Atoi(string(v))

		if err != nil {
			return false
		}

		if (i % 2) == parity {
			num *= 2

			if num > 9 {
				num -= 9
			}
		}

		product += num

	}

	return (product % 10) == 0
}
