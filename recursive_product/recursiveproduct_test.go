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

func TestDivideIntegerInTwo(t *testing.T) {
	cases := []struct {
		name    string
		input   int64
		outputX int64
		outputY int64
	}{
		{
			name:    "should divide into 12345 and 67890",
			input:   1234567890,
			outputX: 12345,
			outputY: 67890,
		},
		{
			name:    "should divide into 1 and 2",
			input:   12,
			outputX: 1,
			outputY: 2,
		},
		{
			name:    "should divide into 12 and 3",
			input:   123,
			outputX: 12,
			outputY: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			x, y := divideIntegerInTwo(c.input)
			assert.Equal(t, c.outputX, x)
			assert.Equal(t, c.outputY, y)
		})
	}
}
