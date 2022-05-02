package main

import (
	"fmt"
)

/** 两次买卖股票的最佳时机
任意一天的状态只能是以下五种状态：
1. 从未买卖过
2. 当前处于首次买入
3. 当前处于首次卖出
4. 当前处于第二次买入
5. 当前处于第二次卖出

由于第一种场景不会产生任何利润，该场景可以不用考虑
最后一天必然处于第5种状态或第2种状态。那么如果prices列表长度为1或者prices一直递减无法获得利润，可以理解为在一天的连续买入，卖出，再买入，卖出

我们将2-5种状态的利润记为 buy1 sell1 buy2 sell2
状态转移方程可以记为：
buy1=max(buy1',-prices[i])  //其中 buy1' 表示前一天已经处于buy1状态，用于往前递推
sell1=max(sell1',buy1'+prices[i])  //其中 sell1' 表示前一天已经处于sell1状态, 也即用于判断是今天卖比较好，还是之前卖比较好
buy2=max(buy2',sell1'-prices[i])
sell2=max(sell2',buy2'+prices[i])

由于同一天买入卖出对收益为0 对结果不会产生影响，上面状态转移方程可以转化为：
buy1=max(buy1,-prices[i]) //可以理解为今天已经buy1状态，或今天刚转移成buy1状态的利润
sell1=max(sell1,buy1+prices[i]) //可以理解为今天已经sell1状态，或今天刚转移成sell1状态的利润
buy2=max(buy2,sell1-prices[i])
sell2=max(sell2,buy2+prices[i])

初始条件
buy1=-prices[0]
sell1=0
buy2=-prices[0]
sell2=0


*/
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return max(max(sell1, sell2), 0)
}

//maxProfit2 原始解法
//buy1=max(buy1',-prices[i])  //其中 buy1' 表示前一天已经处于buy1状态，用于往前递推
//sell1=max(sell1',buy1'+prices[i])  //其中 sell1' 表示前一天已经处于sell1状态, 也即用于判断是今天卖比较好，还是之前卖比较好
//buy2=max(buy2',sell1'-prices[i])
//sell2=max(sell2',buy2'+prices[i])
func maxProfit2(prices []int) int {
	var buy1, sell1, buy2, sell2 int
	buy1Before, sell1Before := -prices[0], 0
	buy2Before, sell2Before := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1Before, -prices[i])
		sell1 = max(sell1Before, buy1Before+prices[i])
		buy2 = max(buy2Before, sell1Before-prices[i])
		sell2 = max(sell2Before, buy2Before+prices[i])

		buy1Before = buy1
		buy2Before = buy2
		sell1Before = sell1
		sell2Before = sell2
	}
	return max(max(sell1, sell2), 0)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}
