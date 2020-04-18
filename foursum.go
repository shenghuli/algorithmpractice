package main

import (
	"fmt"
  "sort"
)

func main(){
  fmt.Println(foursum(([]int{1, 0, -1, 0, -2, 2}, 0))
}


//18. 四数之和
//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
//注意：
//
//答案中不可以包含重复的四元组。
//
//示例：
//
//给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//[-1,  0, 0, 1],
//[-2, -1, 1, 2],
//[-2,  0, 0, 2]
//]

//FourSum
func FourSum(nums []int, target int) [][]int {
	//对原始数组进行排序
	retList := [][]int{}
	if len(nums) < 4 {
		return retList
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] { //去重
				continue
			}
			m := j + 1
			n := len(nums) - 1
			for m < n {
				diff := target - (nums[i] + nums[j] + nums[m] + nums[n])
				if diff == 0 {
					retList = append(retList, []int{nums[i], nums[j], nums[m], nums[n]})
					m++
					n--
					//去重
					for m < n && nums[m] == nums[m-1] {
						m++
					}
					for m < n && nums[n] == nums[n+1] {
						n--
					}
				} else if diff > 0 {
					m++
				} else {
					n--
				}
			}
		}
	}
	return retList
}
