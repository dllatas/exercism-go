package diffsquares

func SquareOfSums(num int) int {
	sums := 0
	for i := 1; i <= num; i++ {
		sums += i
	}
	return sums * sums
}

func SumOfSquares(num int) int {
	sums := 0
	for i := 1; i <= num; i++ {
		sums += (i * i)
	}
	return sums
}

func Difference(num int) int {
	squareos := SquareOfSums(num)
	sumos := SumOfSquares(num)
	return squareos - sumos
}
