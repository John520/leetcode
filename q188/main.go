package main

import (
	"fmt"
	"math"
)

/**
股票可以买卖k次，不能同时参与买卖，也即不能同时买次n只股票，在之后同时卖出的行为
这里需要定义动态规划模型
dp[i][j][0] //表示第i天交易完，j次交易（0表示卖出后），手中没有股票的利润
dp[i][j][1] //表示第i天交易完，j次交易（1表示买入后），手中持有股票的利润

状态转移方程为 ：
dp[i][j][0]=max(dp[i-1][j][0],dp[i-1][j][1]+prices[i])
dp[i][j][1]=max(dp[i-1][j][1],dp[i-1][j-1][0]-prices[i])

初始值
dp[0][0][0]=0
dp[0][1][1]=-prices[0]

返回值
max(dp[n-1][0...k][0]...)

*/
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	}
	//初始化dp
	dp := make([][][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	//初始化利润 -inf
	for j := 0; j < k+1; j++ {
		dp[0][j][1] = math.MinInt32
		dp[0][j][0] = math.MinInt32
	}

	//初始化第0天买入的利润
	dp[0][0][0] = 0
	dp[0][1][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
		}
	}
	res := 0
	for i := 0; i < k+1; i++ {
		res = max(res, dp[len(prices)-1][i][0])
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	//fmt.Println(maxProfit(2,[]int{2,4,1}))
	fmt.Println(maxProfit(2, []int{4, 2, 1, 7}))
}
