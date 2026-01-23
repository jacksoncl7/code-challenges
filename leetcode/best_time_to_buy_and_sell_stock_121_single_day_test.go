// problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	// test table
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 4, 1}, 2},
	}
	for _, test := range tests {
		result := MaxProfit(test.prices)
		if result != test.expected {
			t.Errorf("For examples %v is expected %v but got %v", test.prices, test.expected, result)
		}

		resultBrute := MaxProfitBruteForce(test.prices)
		if resultBrute != test.expected {
			t.Errorf("For examples %v is expected %v but got %v", test.prices, test.expected, resultBrute)
		}
	}
}

func MaxProfitBruteForce(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}

	return maxProfit
}

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	bestProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > bestProfit {
			bestProfit = price - minPrice
		}
	}

	return bestProfit
}
