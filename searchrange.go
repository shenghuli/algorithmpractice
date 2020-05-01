package main

import (
	"fmt"
)

func main(){

  //在排序数组中查找元素的第一个和最后一个位置
	fmt.Println(SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

//34. 在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//如果数组中不存在目标值，返回 [-1, -1]。
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]

//SearchRange 在排序数组中查找元素的第一个和最后一个位置
func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) < 1 {
		return result
	}

	//二分查找最左边的坐标
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	//没找到
	if nums[left] != target {
		return result
	}
	result[0] = left

	//最右边坐标
	left = 0
	right = len(nums) - 1
	for left < right {
		mid := right - (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if nums[right] == target {
		result[1] = right
	}

	return result
}
