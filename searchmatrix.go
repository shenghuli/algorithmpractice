package main

import (
	"fmt"
)

func main(){
  matrix := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	fmt.Println(SearchMatrix(matrix, 1))
}

//74. 搜索二维矩阵
//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//示例 1:
//
//输入:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 3
//输出: true

//SearchMatrix 搜索二维矩阵
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	left := 0
	right := rows*cols - 1
	for left <= right {
		mid := (left + right) / 2
		row := mid / cols
		col := mid % cols
		if target == matrix[row][col] {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
