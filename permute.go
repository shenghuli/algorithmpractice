package main

import (
	"fmt"
)

func main() {
  //全排列
	fmt.Println(Permute([]int{1, 2, 3}))
}

//46. 全排列
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]

//Permute
func Permute(nums []int) [][]int {
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
	for i := 0; i < k; i++ {
		temp := nums[i]
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
