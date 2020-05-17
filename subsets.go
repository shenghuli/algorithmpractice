package main

import (
	"fmt"
)

func main() {
  //78.子集
	fmt.Println(Subsets([]int{1, 2, 3}))
}

//78. 子集
//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//[3],
//[1],
//[2],
//[1,2,3],
//[1,3],
//[2,3],
//[1,2],
//[]
//]

//Subsets 子集
func Subsets(nums []int) [][]int {
	res := &[][]int{}
	backtrack(0, nums, res, []int{})
	return *res
}

//backtrack 回溯
func backtrack(i int, nums []int, res *[][]int, list []int) {
	newList := make([]int, len(list))
	copy(newList, list)
	*res = append(*res, newList)
	for j := i; j < len(nums); j++ {
		list = append(list, nums[j])
		backtrack(j+1, nums, res, list)
		list = list[:len(list)-1]
	}
}
