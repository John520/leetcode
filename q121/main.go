package main

import "fmt"

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	//买卖需在不同一天，若价格列表长度小于1，则利润为0
	if len(prices) <= 1 {
		return 0
	}
	var maxProfit int
	var profit int
	var buyPos int
	var start, end int                 //记录买卖的天数
	diff := make([]int, len(prices)-1) //价格相邻的股票价差
	for i := 0; i < len(prices)-1; i++ {
		diff[i] = prices[i+1] - prices[i]
	}
	for i := 0; i < len(diff); i++ {
		profit += diff[i] //计算持有到现在的利润
		if profit <= 0 {  //累计利润小于等于0，之前可以不进行买入操作
			profit = 0 //重置利润
			buyPos = i //重置下次的买入时间
			continue
		}
		if profit > maxProfit { //之前的理论+今天的价差理论
			start = buyPos
			end = i
			maxProfit = profit
		}
	}
	if maxProfit > 0 {
		fmt.Printf("第%d天买入，第%d天卖出，利润最大%d\n", start+2, end+2, maxProfit) ////i对应的是diff下标，对应到prices需要+1 ，价格是从第一天开始，不是从第0天开始，故需要再+1，所以这里+2
	}
	return maxProfit
}

//maxProfitV2 滑动窗口解法
func maxProfitV2(prices []int) int {
	//买卖需在不同一天，若价格列表长度小于1，则利润为0
	if len(prices) <= 1 {
		return 0
	}
	var maxProfit int
	var profit int
	var left, right int
	var start, end int //记录买卖的天数
	for right < len(prices) {
		profit = prices[right] - prices[left]
		if profit < 0 { //利率小于0
			left = right
			right++
			continue
		}
		if maxProfit < profit {
			maxProfit = profit
			//记录买入和卖出如期
			start = left
			end = right
		}
		right++
	}
	if maxProfit > 0 {
		fmt.Printf("第%d天买入，第%d天卖出，利润最大%d\n", start+1, end+1, maxProfit)
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
	maxProfitV2(prices)
	prices = []int{7, 6, 4, 3, 1}
	maxProfit(prices)
	maxProfitV2(prices)
}
