// 2149. Rearrange Array Elements by Sign

package main

import "fmt"

func main() {
	nums := []int{3, 1, -2, -5, 2, -4}
	solution := rearrangeArray(nums)
	fmt.Println(solution)
}

func rearrangeArray(nums []int) []int {
	var pos []int
	var negs []int
	var solution []int

	for _, num := range nums {
		if num > 0 {
			pos = append(pos, num)
		} else {
			negs = append(negs, num)
		}
	}

	qtPos := 0
	qtNeg := 0

	for index, _ := range nums {

		if index%2 == 0 {
			solution = append(solution, pos[qtPos])
			qtPos++
		} else {
			solution = append(solution, negs[qtNeg])
			qtNeg++
		}

	}

	return solution
}
