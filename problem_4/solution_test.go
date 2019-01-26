package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithMap(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3}, 4},
		{[]int{4, 2, 3}, 1},
		{[]int{1}, 2},
		{[]int{2}, 1},
		{[]int{-2}, 1},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{1, 2, 0}, 3},
	}

	for _, test := range tests {
		desc := fmt.Sprintf("%v to %v", test.input, test.expect)
		t.Run(desc, func(t *testing.T) {
			result := WithMap(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}
