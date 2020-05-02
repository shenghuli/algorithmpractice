package main

import (
	"fmt"
)

func main(){
  //长度最小的子数组
	fmt.Println(MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

//209. 长度最小的子数组
//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。
//
//示例:
//
//输入: s = 7, nums = [2,3,1,2,4,3]
//输出: 2
//解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
//进阶:
//
//如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

//MinSubArrayLen 长度最小的子数组 滑动窗口解法（双指针）
func MinSubArrayLen(s int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum := 0
	left := 0
	result := int(math.MaxInt64)
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= s {
			diff := right - left + 1
			if diff < result {
				result = diff
			}
			sum -= nums[left]
			left++
		}
	}
	if result == int(math.MaxInt64) {
		return 0
	}
	return result
}
