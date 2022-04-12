package main

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	diff := make([]int, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		diff[i] = prices[i+1] - prices[i]
	}
	var sum int
	for _, profit := range diff {
		if profit > 0 {
			sum += profit
		}
	}
	return sum
}
