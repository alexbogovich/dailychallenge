package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductionExceptIndex(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{1, 2, 0}, []int{0, 0, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{-1, 2, -3, 4, -5}, []int{120, -60, 40, -30, 24}},
		{[]int{1, 2, 0, 4, 5}, []int{0, 0, 40, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{0}},
		{[]int{-1}, []int{0}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		desc := fmt.Sprintf("%v to %v", test.input, test.expect)
		t.Run(desc, func(t *testing.T) {
			result := ProductionExceptIndex(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestGetMultilyFromBack(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{6, 3, 1}},
		{[]int{1, 2, 0}, []int{0, 0, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 20, 5, 1}},
		{[]int{1}, []int{0}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		desc := fmt.Sprintf("%v to %v", test.input, test.expect)
		t.Run(desc, func(t *testing.T) {
			result := GetMultiplyFromBack(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}
