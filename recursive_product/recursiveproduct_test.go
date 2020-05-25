package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigits(t *testing.T) {
	cases := []struct {
		name   string
		input  int64
		output int
	}{
		{
			name:   "should return 4",
			input:  1234,
			output: 4,
		},
		{
			name:   "should return 1",
			input:  1,
			output: 1,
		},
		{
			name:   "should return 10",
			input:  1234567890,
			output: 10,
		},
		{
			name:   "should return 10",
			input:  000000,
			output: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o := countDigits(c.input)
			assert.Equal(t, c.output, o)
		})
	}
}
