package main

import (
	"fmt"
)

func main() {
  //全排列 II
	fmt.Println(permuteunique.PermuteUnique([]int{2, 2, 1, 1}))
}

//47. 全排列 II
//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
//示例:
//
//输入: [1,1,2]
//输出:
//[
//[1,1,2],
//[1,2,1],
//[2,1,1]
//]

//PermuteUnique 全排列 II
func PermuteUnique(nums []int) [][]int {
	n := len(nums)
	if n < 1 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{nums}
	}
	retList := [][]int{}
	makePermute(nums, &retList, n, n)
	return retList
}

func makePermute(nums []int, list *[][]int, n, k int) {
	if k == 1 {
		*list = append(*list, nums)
		return
	}
	//任何一个数字都可以作为全排列的最后一位
	uniqDic := map[int]bool{}
	for i := 0; i < k; i++ {
		temp := nums[i]
		//去重
		if _, ok := uniqDic[temp]; ok {
			continue
		} else {
			uniqDic[temp] = true
		}
		nums[i] = nums[k-1]
		nums[k-1] = temp
		//进行神拷贝，裂变
		tmpNums := make([]int, n)
		copy(tmpNums, nums)
		makePermute(tmpNums, list, n, k-1)
		//恢复愿数组
		temp = nums[i]
		nums[i] = nums[k-1]
		nums[k-1] = temp
	}
}
