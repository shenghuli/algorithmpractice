package main

import (
	"fmt"
)

func main(){
  //唯一路径
	fmt.Println(UniqPaths(13, 23))
}

//import "fmt"

//62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//问总共有多少条不同的路径？
//
//
//
//例如，上图是一个7 x 3 的网格。有多少可能的路径？
//
//
//
//示例 1:
//
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右

//UniqPaths 不同路径 组合解法C(m+n-2,m-1),会存在越界问题
func UniqPaths(m int, n int) int {
	var result = 1

	for i := m + n - 2; i >= m; i-- {
		result *= i
	}

	var j = 1
	for i := 1; i <= n-1; i++ {
		j *= i
	}

	return result / j
}

//UniqPaths2 不同的路径动态规划解法
func UniqPaths2(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j] + result[i][j-1]
			}
		}
	}
	return result[m-1][n-1]
}
