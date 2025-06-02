// https://leetcode.com/problems/remove-element/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// val := 2

	result := removeElement(nums, val)
	fmt.Println("RESULT => ", result)
}

func removeElement(nums []int, val int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
