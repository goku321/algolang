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
	}{}

	for _, c := range cases {
		x, y := mergeAndCountInversions(c.inputX, c.inputY)
		assert.Equal(t, x, c.outputSlice)
		assert.Equal(t, y, c.outputCount)
	}
}
