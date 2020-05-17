package main

import (
	"fmt"
)

func main(){
  //77. 组合
	fmt.Println(Combine(4, 2))
}


//77. 组合

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

//Combine 组合
func Combine(n int, k int) [][]int {
	if n < k {
		return [][]int{}
	}
	res := &[][]int{}
	backtrack(1, n, k, res, []int{})
	return *res
}

//backtrack 回溯
func backtrack(i, n, k int, res *[][]int, list []int) {
	if k == 0 {
		//深拷贝
		newList := make([]int, len(list))
		copy(newList, list)
		*res = append(*res, newList)
		return
	}
	for j := i; j <= n; j++ {
		list = append(list, j)
		backtrack(j+1, n, k-1, res, list)
		list = list[:len(list)-1] //恢复至上一个状态
	}
}
