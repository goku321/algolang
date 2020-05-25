package main

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
