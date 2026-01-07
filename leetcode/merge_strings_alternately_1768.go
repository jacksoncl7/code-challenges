// https://leetcode.com/problems/merge-strings-alternately/description

package main

import "fmt"

func main() {
	// input := []int{1, 1, 1, 2, 2, 2, 2, 0, 0, 0, 0, 2, 7}
	input1 := "xyzasc"
	input2 := "a bc888"
	fmt.Println("The alternately vertion of", input2, "and", input1, "is:", mergeAlternately(input1, input2))
}

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		result += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	if i < len(word1) {
		result += word1[i:]
	}
	if j < len(word2) {
		result += word2[j:]
	}

	return result
}
