package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "should return 2",
			input: []int{3, 8, 2, 5, 1, 4, 7, 6},
			want:  2,
		},
		{
			name:  "should return 0",
			input: []int{1, 8, 7, 4, 2},
			want:  0,
		},
		{
			name:  "should return 4",
			input: []int{9, 8, 7, 4, 2},
			want:  4,
		},
		{
			name:  "should return 0",
			input: []int{9},
			want:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := partition(c.input)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestQuickSort(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "sorted slice should be: [1, 2, 3, 4, 5, 6, 7, 8]",
			input: []int{3, 8, 2, 5, 1, 4, 7, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "sorted slice should be: [1, 2]",
			input: []int{2, 1},
			want:  []int{1, 2},
		},
		{
			name:  "sorted slice should be: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			quicksort(c.input)
			assert.Equal(t, c.want, c.input)
		})
	}
}
