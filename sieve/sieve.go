package sieve

func Sieve(limit int) []int {
	var primes []int

	// Bool slice to cache prime numbers,
	// mark all numbers first as primes.
	cache := make([]bool, (limit + 1))
	for i := 2; i <= limit; i++ {
		cache[i] = true
	}

	for i := 2; i <= limit; i++ {
		if cache[i] {
			primes = append(primes, i)

			// Mark composite numbers as non-primes.
			for j := i * i; j <= limit; j += i {
				cache[j] = false
			}
		}
	}

	return primes
}
