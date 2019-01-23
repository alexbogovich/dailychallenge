package main

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsArrayContainsGivenSum(t *testing.T) {
    tests := []struct {
        array            []int
        k                int
        arrayContainsSum bool
    }{
        {[]int{10, 15, 3, 7}, 17, true},
        {[]int{10, 15, 3, 7}, 10, true},
        {[]int{10, 15, 3, 7}, 13, true},
        {[]int{10, 15, 3, 7}, 1, false},
        {[]int{10, 15, 3, 7}, -10, false},
        {[]int{-10, -15, -3, -7}, -17, true},
        {[]int{-10, -15, -3, -7}, -16, false},
        {[]int{}, 1, false},
        {[]int{0, 0}, 0, true},
    }

    for _, test := range tests {
        desc := fmt.Sprintf("%v is sum of some pair in %v", test.k, test.array)
        if !test.arrayContainsSum {
            desc = fmt.Sprintf("%v is not sum of some pair in %v", test.k, test.array)
        }
        t.Run(desc, func(t *testing.T) {
            result := IsArrayContainsGivenSum(test.array, test.k)
            assert.Equal(t, test.arrayContainsSum, result)
        })
    }
}
