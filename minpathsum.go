package main

import (
	"fmt"
)

func main(){
	//64. 最小路径和
	//fmt.Println(MinPathSum([][]int{
	//	[]int{1, 3, 1},
	//	[]int{1, 5, 1},
	//	[]int{4, 2, 1},
	//}))
}


//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//示例:
//
//输入:
//[
//  [1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。

//MinPathSum 最小路径和
func MinPathSum(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	minPath := make([][]int, rows) //记录每个节点的最小路径和
	for i := 0; i < rows; i++ {
		minPath[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j == 0 {
				minPath[i][j] = grid[i][j]
			} else {
				if i == 0 {
					minPath[i][j] = minPath[i][j-1] + grid[i][j]
				} else if j == 0 {
					minPath[i][j] = minPath[i-1][j] + grid[i][j]
				} else if minPath[i-1][j] < minPath[i][j-1] {
					minPath[i][j] = minPath[i-1][j] + grid[i][j]
				} else {
					minPath[i][j] = minPath[i][j-1] + grid[i][j]
				}
			}
		}
	}

	return minPath[rows-1][cols-1]
}
