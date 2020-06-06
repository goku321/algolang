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

func TestCountInversions(t *testing.T) {
	cases := []struct {
		name      string
		input     []int
		wantSlice []int
		wantCount int
	}{
		{
			name:      "should return [1, 2, 3, 4] and 0",
			input:     []int{1, 2, 3, 4},
			wantSlice: []int{1, 2, 3, 4},
			wantCount: 0,
		},
		{
			name:      "should return [1, 2, 3, 4] and 2",
			input:     []int{2, 1, 4, 3},
			wantSlice: []int{1, 2, 3, 4},
			wantCount: 2,
		},
		{
			name:      "should return [1, 2, 3, 4] and 6",
			input:     []int{4, 3, 2, 1},
			wantSlice: []int{1, 2, 3, 4},
			wantCount: 6,
		},
		{
			name:      "should return [1, 2, 3, 4, 5] and 10",
			input:     []int{5, 4, 3, 2, 1},
			wantSlice: []int{1, 2, 3, 4, 5},
			wantCount: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o, count := countInversions(c.input)
			assert.Equal(t, c.wantSlice, o)
			assert.Equal(t, c.wantCount, count)
		})
	}
}
