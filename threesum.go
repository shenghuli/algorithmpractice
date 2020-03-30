package main

import (
	"fmt"
  "sort"
 )
 
 func main() {
    //三数之和
	  list := []int{0, 0, -1, 0, 1, 2, -1, -4}
	  fmt.Println(threesum.ThreeSum(list))
 }
 
 //三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//参考解法：https://blog.csdn.net/qq_40817333/article/details/100593420?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522158542994219724846461375%2522%252C%2522scm%2522%253A%252220140713.130056874..%2522%257D&request_id=158542994219724846461375&biz_id=0&utm_source=distribute.pc_search_result.none-task

// ThreeSum 三数之和
func ThreeSum(nums []int) [][]int {
	retList := [][]int{}
	if len(nums) < 3 {
		return retList
	}
	//排序
	sort.Ints(nums)
	var target int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target = nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r]+target < 0 {
				l++
			} else if nums[l]+nums[r]+target > 0 {
				r--
			} else { //和为0
				retList = append(retList, []int{target, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return retList
}
