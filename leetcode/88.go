// 88. Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/description/
// challenge O(n+m) =)
package main

// func main() {
// 	// nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
// 	// nums2 := []int{1, 2, 2}

// 	nums1 := []int{1, 2, 3, 0}
// 	nums2 := []int{1}

// 	result := merge(nums1, 3, nums2, 1)
// 	fmt.Println("RESULT => ", result)
// }

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var index1, index2 int
	aux := []int{}

	if m > 0 && n > 0 {
		for {
			// fmt.Println(aux, index1, index2)
			if nums1[index1] == nums2[index2] {
				aux = append(aux, nums1[index1], nums2[index2])
				index2++
				index1++
			} else {
				if nums1[index1] > nums2[index2] {
					aux = append(aux, nums2[index2])
					index2++
				} else {
					aux = append(aux, nums1[index1])
					index1++
				}
			}

			if index1 > m-1 {
				aux = append(aux, nums2[index2:n]...)
				break
			}
			if index2 > n-1 {
				aux = append(aux, nums1[index1:m]...)
				break
			}

			if index1 > m && index2 > n {
				break
			}

		}
	}

	if m == 0 {
		for i, _ := range nums2 {
			nums1[i] = nums2[i]
		}
	} else {
		for i, _ := range aux {
			nums1[i] = aux[i]
		}
	}

	return nums1
}
