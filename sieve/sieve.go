package sieve

func Sieve(limit int) []int {
	primes := make([]int, 0, (limit / 2))

	// Bool slice to cache composite numbers.
	composite := make([]bool, (limit + 1))

	for i := 2; i <= limit; i++ {
		if !composite[i] {
			primes = append(primes, i)

			// Mark all multiples of i as composite.
			for j := i * i; j <= limit; j += i {
				composite[j] = true
			}
		}
	}

	return primes
}
