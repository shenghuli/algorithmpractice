package main

import (
	"fmt"
)

func main(){
  //73. 矩阵置零
	matrix := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	SetZeroes(matrix)
	PrintMatrix(matrix)
}

//73. 矩阵置零
//给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
//
//示例 1:
//
//输入:
//[
//[1,1,1],
//[1,0,1],
//[1,1,1]
//]
//输出:
//[
//[1,0,1],
//[0,0,0],
//[1,0,1]
//]

//SetZeroes 矩阵置零
func SetZeroes(matrix [][]int) {
	fr, fc := false, false //标记首行，首列是否为0
	if len(matrix) < 1 {
		return
	}
	//将值为0的元素，下沉至首行、首列
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					fr = true
				}
				if j == 0 {
					fc = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if fr {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if fc {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

//PrintMatrix 打印矩阵
func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
