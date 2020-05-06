package main

import (
	"fmt"
  "math"
)

func main(){
  //最大子序和
	fmt.Println(MaxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}


//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

//解法1，大体思路：记录一个0～i的和
//记录一个0～i之前的最小小和
//两者做差就是最大子序和

//MaxSubArray 最大子序和
func MaxSubArray(nums []int) int {
	allSum := 0
	minSum := 0
	res := int(math.MinInt64)
	for i := 0; i < len(nums); i++ {
		allSum += nums[i]
		if allSum-minSum > res {
			res = allSum - minSum
		}
		if allSum < minSum {
			minSum = allSum
		}
	}
	return res
}

//MaxSubArray2 动态规划解法
func MaxSubArray2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	result := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] { //前面的连续和大于当前值
			dp[i] = dp[i-1] + nums[i]
		} else { //当前值比较大
			dp[i] = nums[i]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}
