package main

// merges two sorted slices and count inversions.
func mergeAndCountInversions(x []int, y []int) ([]int, int) {
	lenX := len(x)
	lenY := len(y)

	c := make([]int, lenX+lenY)
	inversions := 0
	i, j := 0, 0

	for k := range c {
		if i < lenX && j < lenY && x[i] <= y[j] {
			c[k] = x[i]
			i++
		} else if i < lenX && j < lenY {
			c[k] = y[j]
			inversions += lenX - i
			j++
		} else if i < lenX {
			c[k] = x[i]
			i++
		} else {
			c[k] = y[j]
			j++
		}
	}

	return c, inversions
}

func countInversions(x []int) ([]int, int) {
	lenX := len(x)
	if lenX == 1 {
		return x, 0
	}

	mid := lenX / 2
	sortedLeft, leftCount := countInversions(x[0:mid])
	sortedRight, rightCount := countInversions(x[mid:lenX])
	mergedSlice, splitCount := mergeAndCountInversions(sortedLeft, sortedRight)

	return mergedSlice, (leftCount + rightCount + splitCount)
}
