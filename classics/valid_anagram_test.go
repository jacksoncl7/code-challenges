package classics

import "testing"

func TestIsAnagram(t *testing.T) {
	// test table
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"danger", "garden", true},
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"banana", "cabana", false},
		{"", "", true},
		{" ", " ", true},
	}
	for _, test := range tests {
		result := isAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("For examples %s and %s is expected %v but got %v", test.s, test.t, test.expected, result)
		}
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterCount := make(map[rune]int, len(s))
	for _, char := range s {
		letterCount[char]++
	}

	for _, char := range t {
		letterCount[char]--
		if letterCount[char] < 0 {
			return false
		}
	}

	for _, count := range letterCount {
		if count != 0 {
			return false
		}
	}

	return true
}
