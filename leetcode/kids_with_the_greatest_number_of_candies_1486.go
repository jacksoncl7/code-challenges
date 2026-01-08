//https://leetcode.com/problems/kids-with-the-greatest-number-of-candies

package main

import "fmt"

func main() {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	fmt.Println("Kids with the greatest number of candies:", kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxCandies := 0

	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}
