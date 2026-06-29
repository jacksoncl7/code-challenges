// problem: https://leetcode.com/problems/greatest-common-divisor-of-strings/description

package main

import (
	"testing"
)

func TestGreatestCommonDivisorOfStrings(t *testing.T) {
	tests := []struct {
		inputFirst     string
		inputSecond    string
		expectedOutput string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "AB", "AB"},
		{"LEET", "LEET", "LEET"},
		{"CODE", "CODE", "CODE"},
	}

	for _, test := range tests {
		result := greatestCommonDivisorOfStrings(test.inputFirst, test.inputSecond)
		if result != test.expectedOutput {
			t.Errorf("For examples %v and %v is expected %v but got %v", test.inputFirst, test.inputSecond, test.expectedOutput, result)
		}
	}
}

func greatestCommonDivisorOfStrings(inputFirst, inputSecond string) string {
	if inputFirst+inputSecond != inputSecond+inputFirst {
		return ""
	}

	lenFirst, lenSecond := len(inputFirst), len(inputSecond)
	for lenSecond > 0 {
		lenFirst, lenSecond = lenSecond, lenFirst%lenSecond
	}

	return inputFirst[:lenFirst]
}
