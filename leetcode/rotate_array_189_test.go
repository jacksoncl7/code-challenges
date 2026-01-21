//https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	// test table
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{-1, -100, 3, 99}, 1, []int{99, -1, -100, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{5, 6, 7, 8, 1, 2, 3, 4}},
		{[]int{7}, 1, []int{7}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2}, 7, []int{2, 1}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
	}
	for _, test := range tests {

		// Rotate(test.input, test.k)
		RotateOptimal(test.input, test.k)

		for i, v := range test.input {
			if v != test.expected[i] {
				t.Errorf("For examples %v and %d is expected %v but got %v", test.input, test.k, test.expected, test.input)
				break
			}
		}
	}
}

// brute force O(n*k)
func Rotate(input []int, k int) {
	lenInput := len(input)
	k = k % lenInput

	if k == 0 {
		return
	}

	if lenInput < 2 {
		return
	}

	for range k {
		lastElement := input[lenInput-1]
		for i := lenInput - 1; i > 0; i-- {
			input[i] = input[i-1]
		}
		input[0] = lastElement
	}
}

// optimal O(n) time and O(1) space
func RotateOptimal(input []int, k int) {
	lenInput := len(input)
	k = k % lenInput

	if k == 0 {
		return
	}

	if lenInput < 2 {
		return
	}

	reverse(input, 0, lenInput-1)
	reverse(input, 0, k-1)
	reverse(input, k, lenInput-1)
}

func reverse(input []int, start int, end int) {
	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}
}
