package main

import (
	"fmt"
)

func main(){
  //缺失的第一个正数
	fmt.Println(FirstMissingPositive([]int{3, 4, -1, 1}))
}

//41. 缺失的第一个正数
//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//
//
//
//示例 1:
//
//输入: [1,2,0]
//输出: 3

//参考：https://zhuanlan.zhihu.com/p/115945810

//大概思路，通过一遍循环，将每个数字放在其num-1的位置
//这样从左到右遍历一遍，第一个不等于其index+1的位置对应的数，就是第一个缺失的数

func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] < 1 { //数字不是正整数
			i++
		} else if nums[i] > n { //数字超出数组右边界
			i++
		} else if nums[i] == i+1 { //数字刚好等于i+1，无需交换
			i++
		} else if nums[i] != nums[nums[i]-1] { //将nums[i]放置到其对应的位置上
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		} else { //重复元素，将其置为1～n区间外的数
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	//默认返回最大值n+1
	return n + 1
}
