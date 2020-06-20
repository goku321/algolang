package main

// partitions the slice around the pivot x[0].
//
// returns the final position of pivot.
func partition(x []int) int {
	if len(x) < 3 {

	}

	pivot := x[0]
	i := 1

	for j := i; j < len(x); j++ {
		if x[j] < pivot {
			temp := x[j]
			x[j] = x[i]
			x[i] = temp
			i++
		}
	}

	// Place the pivot at correct position.
	temp := x[i-1]
	x[i-1] = pivot
	x[0] = temp

	return i - 1
}
