package main

import (
	"fmt"
)


func main() {

	//股票最大收益III
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(MaxProfit(nums))

}


//买卖股票的最佳时机 III
//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//示例 1:
//
//输入: [3,3,5,0,0,3,1,4]
//输出: 6
//解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。

// MaxProfit 买卖股票的最佳时机 III
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	//第i天之前卖出获得的最大利润
	profitLeft := make([]int, len(prices))
	minPrice := prices[0]
	profitLeft[0] = 0
	for i := 1; i < len(prices); i++ {
		if minPrice-prices[i] > 0 {
			minPrice = prices[i]
		}
		if profitLeft[i-1] > prices[i]-minPrice {
			profitLeft[i] = profitLeft[i-1]
		} else {
			profitLeft[i] = prices[i] - minPrice
		}
	}
	//第i天之后买入获得的最大利润
	profitRight := make([]int, len(prices))
	profitRight[len(prices)-1] = 0
	maxPrice := prices[len(prices)-1]
	for i := len(prices) - 2; i > 0; i-- {
		if maxPrice < prices[i] {
			maxPrice = prices[i]
		}
		if profitRight[i-1] > maxPrice-prices[i] {
			profitRight[i] = profitRight[i-1]
		} else {
			profitRight[i] = maxPrice - prices[i]
		}
	}

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if maxProfit < profitLeft[i]+profitRight[i] {
			maxProfit = profitLeft[i] + profitRight[i]
		}
	}
	return maxProfit
}
