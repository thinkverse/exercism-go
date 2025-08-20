package luhn

// Valid check is a string is a valid LUHN number.
//
// Benchmarks done with Core i5 760 2.80GHz
//
// Operations  Time (ns/op)
// 271239      3730
//
// 279519      3726
//
// 1189321     1021
//
// 3331646     350
func Valid(id string) bool {
	var digits []int

	for _, v := range id {
		if v == ' ' {
			continue
		}

		if !('0' <= v && v <= '9') {
			return false
		}

		digits = append(digits, int((v - '0')))
	}

	length := len(digits)

	if length <= 1 {
		return false
	}

	parity := (length - 2) % 2

	var product int

	for i, v := range digits {
		if (i % 2) == parity {
			v *= 2

			if v > 9 {
				v -= 9
			}
		}

		product += v

	}

	return (product % 10) == 0
}
