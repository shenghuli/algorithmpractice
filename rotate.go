package main

import (
	"fmt"
)

func main(){
	//旋转图像
	m := [][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}}
	Rotate(m)
	fmt.Println(m)
}


//48. 旋转图像
//给定一个 n × n 的二维矩阵表示一个图像。
//
//将图像顺时针旋转 90 度。
//
//说明：
//
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
//示例 1:
//
//给定 matrix =
//[
//[1,2,3],
//[4,5,6],
//[7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[7,4,1],
//[8,5,2],
//[9,6,3]
//]

func Rotate(matrix [][]int) {
	n := len(matrix) //nxn方阵
	if n < 1 {
		return
	}
	if len(matrix[0]) != n { //非方阵
		return
	}
	for i := 0; i < n/2; i++ {
		//旋转上横行,旋转n-1一个数字
		for l := i; l < n-i-1; l++ {
			//构造右边
			tmp1 := matrix[l][n-i-1]
			matrix[l][n-i-1] = matrix[i][l]
			//构造下边
			tmp2 := matrix[n-i-1][n-l-1]
			matrix[n-i-1][n-l-1] = tmp1
			//构造左边
			tmp1 = matrix[n-l-1][i]
			matrix[n-l-1][i] = tmp2
			//构造上边
			matrix[i][l] = tmp1
		}
	}
}

