package main

import (
	"fmt"
)

func main() {
	//股票最大收益II
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(MaxProfit(nums))
}


//买卖股票的最佳时机 II
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	val := 0
	//前一天的价钱
	for i := 1; i < len(prices); i++ {
		currentDiff := prices[i] - prices[i-1]
		if currentDiff > 0 { //卖出
			val += currentDiff
		}
		//买入，买入不赚钱
	}
	return val
}
