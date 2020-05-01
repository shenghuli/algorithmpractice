package main

import (
	"fmt"
)

func main(){
  //组合总和
	fmt.Println(CombinationSum([]int{2, 3, 6, 7}, 7))
}

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1:
//
//输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//[7],
//[2,2,3]
//]

//思路：https://zhuanlan.zhihu.com/p/87156985
//跟题其实是深度优先搜索的变体

//CombinationSum 组合总和,递归求解
func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return [][]int{}
	}
	result := [][]int{}
	list := []int{}
	commbinationSumCore(&result, list, &candidates, target, 0)
	return result
}

func commbinationSumCore(result *[][]int, list []int, candidates *[]int, remain, start int) {
	//该path无效，直接回溯
	if remain < 0 {
		return
	}

	//该list有效
	if remain == 0 {
		//list为指针类型，需要生成新地址
		newlist := make([]int, len(list))
		copy(newlist, list)
		*result = append(*result, newlist)
	}
	//递归
	for i := start; i < len(*candidates); i++ {
		newRemain := remain - (*candidates)[i]
		list = append(list, (*candidates)[i]) //添加
		commbinationSumCore(result, list, candidates, newRemain, i)
		list = list[:len(list)-1] //回溯
	}
}
