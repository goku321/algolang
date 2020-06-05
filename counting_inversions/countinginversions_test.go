package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAndCountInversions(t *testing.T) {
	cases := []struct {
		name        string
		inputX      []int
		inputY      []int
		outputSlice []int
		outputCount int
	}{
		{
			name:        "should return [1, 2, 3, 4, 5, 6]",
			inputX:      []int{2, 4, 6},
			inputY:      []int{1, 3, 5},
			outputSlice: []int{1, 2, 3, 4, 5, 6},
			outputCount: 6,
		},
		{
			name:        "should return [1]",
			inputX:      []int{1},
			inputY:      []int{},
			outputSlice: []int{1},
			outputCount: 0,
		},
		{
			name:        "should return []",
			inputX:      []int{},
			inputY:      []int{},
			outputSlice: []int{},
			outputCount: 0,
		},
		{
			name:        "should return [1, 1, 2, 2, 3, 3]",
			inputX:      []int{1, 2, 3},
			inputY:      []int{1, 2, 3},
			outputSlice: []int{1, 1, 2, 2, 3, 3},
			outputCount: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			x, y := mergeAndCountInversions(c.inputX, c.inputY)
			assert.Equal(t, c.outputSlice, x)
			assert.Equal(t, c.outputCount, y)
		})

	}
}
