package main

import (
	"fmt"
)

func main(){
  //搜索旋转排序数组
	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}


//33. 搜索旋转排序数组

//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4

func Search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//此时分割点为left,然后分别从每个区间找
	spit_point := left
	fmt.Println(spit_point)
	left = 0
	right = spit_point - 1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//从右边区间找
	left = spit_point
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
