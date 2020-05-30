package main

import "math"

// Count number of digits in a positive integer.
// Returns 1 for 0.
func countDigits(x int64) int {
	cnt := 0

	if x == 0 {
		return 1
	}

	for x > 0 {
		x /= 10
		cnt++
	}

	return cnt
}

func splitIntegerInTwo(x int64) (a, b int64) {
	n := countDigits(x)

	firstHalf := x / int64(math.Pow10(n/2))
	secondHalf := x % int64(math.Pow10(n/2))

	return firstHalf, secondHalf
}

// Calculates product b/w x and y.
// x and y have even and equal number of digits.
// func recursiveProduct(x, y int64) int64 {
// 	// Base case
// 	if countDigits(x) == 1 {
// 		return x * y
// 	}

// 	cntX := countDigits(x)
// 	cntY := countDigits(y)
// }
