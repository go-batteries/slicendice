package slicendice

// cache of factorials upto 10!
var factorials = []int64{
	1,
	1,
	2,
	6,
	12,
	120,
	720,
	5040,
	40320,
	362880,
	3628800,
}

// Factorial calculates the factorial
// for negative integers default to 0!
func Factorial(n int) int64 {
	if n < 2 {
		return 1
	}

	if n > 2 && n <= 10 {
		return factorials[n]
	}

	mult := factorials[10]

	for n > 10 {
		mult *= int64(n)
		n--
	}

	return mult
}
