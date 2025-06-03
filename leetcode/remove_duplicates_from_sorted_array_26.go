// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

// func main() {
// 	// Input: nums = [1,1,2]
// 	// Output: 2, nums = [1,2,_]
// 	// input := []int{1, 1, 2}
// 	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}
// 	//OUT   nums: [0  1  2  3  4  2  2  3  3  4  4]
// 	fmt.Println("RESULT =>", removeDuplicatesPart1(input))
// }

func removeDuplicatesPart1(nums []int) int {
	counter := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
