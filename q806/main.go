package main

import "fmt"

//https://leetcode-cn.com/problems/number-of-lines-to-write-string/

func numberOfLines(widths []int, s string) []int {
	ans := make([]int, 0, 2)
	rows := 1
	var rowSum int
	for _, c := range s {
		if rowSum+widths[c-'a'] > 100 {
			rows++
			rowSum = widths[c-'a']
		} else {
			rowSum += widths[c-'a']
		}
	}
	ans = append(ans, rows)
	ans = append(ans, rowSum)
	return ans
}

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	//fmt.Println(numberOfLines(widths, s))
	widths = []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s = "bbbcccdddaaa"
	fmt.Println(numberOfLines(widths, s))
}
