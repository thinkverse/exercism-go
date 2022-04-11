package diffsquares

func SquareOfSum(n int) int {
	term := ((n * (n + 1)) / 2)

	return term * term
}

func SumOfSquares(n int) int {
	term := (n * (n + 1) * ((2 * n) + 1))

	return term / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
