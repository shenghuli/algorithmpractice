package main

import (
	"fmt"
  "sort"
)

func main(){
	//组合总和2
  fmt.Println(CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}


//思路参见CombinationSum

//CombinationSum2 组合总和,递归求解
func CombinationSum2(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := [][]int{}
	combinationSum2Core(&result, []int{}, candidates, target, 0)
	return result
}

func combinationSum2Core(result *[][]int, list []int, candidates []int, remain, start int) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		newList := make([]int, len(list))
		copy(newList, list)
		*result = append(*result, newList)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > remain {
			return
		}
		list = append(list, candidates[i])
		combinationSum2Core(result, list, candidates, remain-candidates[i], i+1)
		list = list[:len(list)-1]
		//去重
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}
}
