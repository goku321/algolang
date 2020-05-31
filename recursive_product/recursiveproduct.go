package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

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

// Makes strings equal in length by padding zeroes
// to the left of the shorter string.
func makeStringsEqualInLength(x, y string) (string, string) {
	diff := len(x) - len(y)
	if diff == 0 {
		return x, y
	}

	if diff > 0 {
		y = fmt.Sprintf("%0*s", len(x), y)
		return x, y
	}

	x = fmt.Sprintf("%0*s", len(y), x)
	return x, y
}

// Adds two integers represented as strings.
// Input strings must of same length.
func addIntegerStrings(x, y string) string {
	strLen := len(x)
	resStr := ""
	c := 0

	for i := strLen - 1; i >= 0; i-- {
		a, err := strconv.Atoi(string(x[i]))
		if err != nil {
			log.Printf("invalid input: %v", err)
		}

		b, err := strconv.Atoi(string(y[i]))
		if err != nil {
			log.Printf("invalid input: %v", err)
		}

		sum := a + b + c
		resStr = strconv.Itoa(sum%10) + resStr
		c = sum / 10
	}

	if c > 0 {
		resStr = strconv.Itoa(c) + resStr
	}

	return resStr
}

// Calculates product b/w x and y.
// x and y have even and equal number of digits.
func recursiveProduct(x, y int64) int64 {
	// Base case
	if countDigits(x) == 1 {
		return x * y
	}

	cntX := countDigits(x)

	a, b := splitIntegerInTwo(x)
	c, d := splitIntegerInTwo(y)

	product1 := int64(math.Pow10(cntX)) * recursiveProduct(a, c)
	product2 := int64(math.Pow10(cntX/2)) * (recursiveProduct(a, d) + recursiveProduct(b, c))
	product3 := recursiveProduct(b, d)

	return product1 + product2 + product3
}
