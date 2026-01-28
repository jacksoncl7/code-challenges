// problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"testing"
)

func TestMaxProfitMultipleBuys(t *testing.T) {
	// test table
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		result := MaxProfitMultipleBuys(test.prices)
		if result != test.expected {
			t.Errorf("For examples %v is expected %v but got %v", test.prices, test.expected, result)
		}
	}
}

func MaxProfitMultipleBuys(prices []int) int {
	minPrice := prices[0]
	bestProfit := 0
	maxProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > bestProfit {
			maxProfit += price - minPrice
			minPrice = price
			bestProfit = 0
		}
	}

	return maxProfit
}
