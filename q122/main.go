package main

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

//贪心解法
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

//动态规划

//考虑到「不能同时参与多笔交易」，因此每天交易结束后只可能存在手里有一支股票或者没有股票的状态。
//
//定义状态 dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润
//dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润（i 从 0 开始）。
//
//考虑 \textit{dp}[i][0]dp[i][0] 的转移方程，如果这一天交易完后手里没有股票，那么可能的转移状态为前一天已经没有股票，即 \textit{dp}[i-1][0]dp[i−1][0]，或者前一天结束的时候手里持有一支股票，即 \textit{dp}[i-1][1]dp[i−1][1]，这时候我们要将其卖出，并获得 \textit{prices}[i]prices[i] 的收益。因此为了收益最大化，我们列出如下的转移方程：
//
//\textit{dp}[i][0]=\max\{\textit{dp}[i-1][0],\textit{dp}[i-1][1]+\textit{prices}[i]\}
//dp[i][0]=max{dp[i−1][0],dp[i−1][1]+prices[i]}
//~
//再来考虑 \textit{dp}[i][1]dp[i][1]，按照同样的方式考虑转移状态，那么可能的转移状态为前一天已经持有一支股票，即 \textit{dp}[i-1][1]dp[i−1][1]，或者前一天结束时还没有股票，即 \textit{dp}[i-1][0]dp[i−1][0]，这时候我们要将其买入，并减少 \textit{prices}[i]prices[i] 的收益。可以列出如下的转移方程：
//
//\textit{dp}[i][1]=\max\{\textit{dp}[i-1][1],\textit{dp}[i-1][0]-\textit{prices}[i]\}
//dp[i][1]=max{dp[i−1][1],dp[i−1][0]−prices[i]}
//
//对于初始状态，根据状态定义我们可以知道第 00 天交易结束的时候 \textit{dp}[0][0]=0dp[0][0]=0，\textit{dp}[0][1]=-\textit{prices}[0]dp[0][1]=−prices[0]。
//
//因此，我们只要从前往后依次计算状态即可。由于全部交易结束后，持有股票的收益一定低于不持有股票的收益，因此这时候 \textit{dp}[n-1][0]dp[n−1][0] 的收益必然是大于 \textit{dp}[n-1][1]dp[n−1][1] 的，最后的答案即为 \textit{dp}[n-1][0]dp[n−1][0]。

//func maxProfit(prices []int) int {
//    n := len(prices)
//    dp := make([][2]int, n)
//    dp[0][1] = -prices[0]
//    for i := 1; i < n; i++ {
//        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
//        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
//    }
//    return dp[n-1][0]
//}
//
//func max(a, b int) int {
//    if a > b {
//        return a
//    }
//    return b
//}
