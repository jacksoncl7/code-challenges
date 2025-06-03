// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description
package main

// func main() {
// 	input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3} // output 1,1,2,2,3
// 	fmt.Println("RESULT =>", removeDuplicates(input))
// }

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	counter := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[counter-2] {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
