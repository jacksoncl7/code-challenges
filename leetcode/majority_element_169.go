// https://leetcode.com/problems/majority-element/description
package main

// func main() {
// 	// input := []int{1, 1, 1, 2, 2, 2, 2, 0, 0, 0, 0, 2, 7}
// 	input := []int{3, 2, 3}
// 	fmt.Println("Majority Element in ", input, " is ", majorityElement(input))
// }

func majorityElement(nums []int) int {
	aux := map[int]int{}

	for i := 0; i < len(nums); i++ {
		aux[nums[i]]++
	}
	// fmt.Println("Map => ", aux)
	majorityElement := nums[0]
	biggerPosition := 0

	for i, n := range aux {
		if n > biggerPosition {
			biggerPosition = n
			majorityElement = i
		}
	}

	return majorityElement
}
