package differenceofsquares

func SquareOfSum(n int) int {
	sum := 0

	for n > 0 {
		sum += n
		n--
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0

	for n > 0 {
		sum += n * n
		n--
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
