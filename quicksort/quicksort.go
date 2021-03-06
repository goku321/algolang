package main

// partitions the slice around the pivot x[0].
//
// returns the final position of pivot.
func partition(x []int) int {
	if len(x) > 1 {
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
	return 0
}

// sorts x.
func quicksort(x []int) {
	lenX := len(x)
	if lenX > 1 {
		partitionIndex := partition(x)
		quicksort(x[:(partitionIndex + 1)])
		quicksort(x[partitionIndex+1:])
	}
}
