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
