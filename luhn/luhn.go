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
//
// 10332144    112
//
// Thank you to @thomas-holmes[^1], looking at that iteration made my next iteration even faster.
//
// [^1]: https://exercism.org/tracks/go/exercises/luhn/solutions/thomas-holmes
func Valid(id string) bool {
	digits := make([]byte, 0, len(id))

	for i := 0; i < len(id); i++ {
		if id[i] == ' ' {
			continue
		}

		if !('0' <= id[i] && id[i] <= '9') {
			return false
		}

		digits = append(digits, byte((id[i] - '0')))
	}

	length := len(digits)

	if length <= 1 {
		return false
	}

	parity := (length - 2) % 2

	var product int

	for i := 0; i < length; i++ {
		num := int(digits[i])

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
