// 4. Median of Two Sorted Arrays

package main

// func main() {
// 	nums1 := []int{1, 2} // 1 2 3 4 5 5 6
// 	nums2 := []int{3, 4}
// 	solution := findMedianSortedArrays(nums1, nums2)
// 	fmt.Println(solution) // assert 2.5
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float32 {
	lenNum1 := len(nums1)
	lenNum2 := len(nums2)
	var floatingPartialResult float32

	if lenNum1%2 == 0 {
		floatingPartialResult = float32(nums1[lenNum1/2]+nums1[(lenNum1/2)+1]) / 2.0
	} else {
		floatingPartialResult = float32(nums1[(lenNum1/2)+1])
	}

	if lenNum2%2 == 0 {
		particial2 := float32(nums2[lenNum2/2]+nums2[(lenNum2/2)+1]) / 2.0
		floatingPartialResult = (floatingPartialResult + particial2) / 2.0
	} else {
		floatingPartialResult = float32(nums2[(lenNum2/2)+1])
	}

	return floatingPartialResult
}
