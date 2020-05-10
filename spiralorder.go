package main

import (
	"fmt"
)

func main(){
  //螺旋矩阵
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{10, 11, 12},
		[]int{13, 14, 15},
	}
	fmt.Println(SpiralOrder(matrix))
}

//54. 螺旋矩阵
//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//示例 1:
//
//输入:
//[
//[ 1, 2, 3 ],
//[ 4, 5, 6 ],
//[ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]
//示例 2:

//SpiralOrder 剑指offer第29题
func SpiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}
	rows := len(matrix)
	cols := len(matrix[0])
	circle := 0
	for 2*circle < rows && 2*circle < cols { //此处不能用circle/2，会少循环一层
		//上边,肯定有
		for i := circle; i < cols-circle; i++ {
			result = append(result, matrix[circle][i])
		}
		//右边，不一定有
		if rows-circle*2 > 1 {
			for i := circle + 1; i < rows-circle; i++ {
				result = append(result, matrix[i][cols-circle-1])
			}
		}

		//下边,不一定有
		if rows-2*circle > 1 {
			for i := cols - circle - 2; i >= circle; i-- {
				result = append(result, matrix[rows-circle-1][i])
			}
		}

		//左边,不一定有
		if cols-2*circle > 1 {
			for i := rows - circle - 2; i > circle; i-- {
				result = append(result, matrix[i][circle])
			}
		}

		circle++
	}
	return result
}
