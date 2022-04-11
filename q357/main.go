package main

//https://leetcode-cn.com/problems/count-numbers-with-unique-digits/
//当n=0 时 0 <= x < 1 ,只有一种选择 ，即0
//当n=1 时 0 <= x <10, 有0~9，10种可能
//当n=2 时 0 <= x <= 100, 其中x为1位数有上述的10种可能，当x为两位数时，十位有9种可能，个位有9种可能（本来个位有0-9十种可能，题意要求各位数字不同，故有9中可能）
//当n=3 时 0 <= x <= 1000,当x为2位和1位时，同上。当x为三位数时，百位有9种可能，十位有9种可能，各位有8中可能，也即全排列
//故当n>2时，当x有n位时， 9*(9*8*...) 全排列组合
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	//全排列 9* A(9)(n-1)
	ans := 9
	for i := 0; i < n-1; i++ {
		ans *= (9 - i)
	}
	return ans + countNumbersWithUniqueDigits(n-1)
}

//官方解答
//func countNumbersWithUniqueDigits(n int) int {
//	if n == 0 {
//		return 1
//	}
//	if n == 1 {
//		return 10
//	}
//	ans, cur := 10, 9
//	for i := 0; i < n-1; i++ {
//		cur *= 9 - i
//		ans += cur
//	}
//	return ans
//}
