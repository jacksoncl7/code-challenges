//https://leetcode.com/problems/can-place-flowers/description

package main

// import "fmt"

// // func main() {
// // 	flowerbed := []int{0, 0, 0, 0, 0, 0, 0}
// // 	n := 4
// // 	fmt.Println("Can place flowers:", canPlaceFlowers(flowerbed, n))
// // }

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbedLen := len(flowerbed)

	if flowerbed[0] == 0 {
		if flowerbedLen == 1 || flowerbed[1] == 0 {
			flowerbed[0] = 1
			n--
		}
	}

	for i := 1; i < flowerbedLen-1; i++ {
		if flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}

		if n == 0 {
			return true
		}
	}

	if flowerbed[flowerbedLen-1] == 0 && flowerbed[flowerbedLen-2] == 0 {
		flowerbed[flowerbedLen-1] = 1
		n--
	}

	if n > 0 {
		return false
	}
	return true
}
