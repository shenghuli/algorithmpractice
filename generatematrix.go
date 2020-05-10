package main

import (
	"fmt"
)

func main(){
  //螺旋矩阵 II
	fmt.Println(generatematrix.GenerateMatrix(3))
}

//给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
//示例:
//
//输入: 3
//输出:
//[
//[ 1, 2, 3 ],
//[ 8, 9, 4 ],
//[ 7, 6, 5 ]
//]

func GenerateMatrix(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	circle := 0
	start := 1
	for circle*2 < n {
		//打印上边
		for i := circle; i < n-circle; i++ {
			result[circle][i] = start
			start++
		}

		if circle*2 < n {
			//打印右边
			for i := circle + 1; i < n-circle; i++ {
				result[i][n-circle-1] = start
				start++
			}

			//打印下边
			for i := n - circle - 2; i >= circle; i-- {
				result[n-circle-1][i] = start
				start++
			}

			//打印左边
			for i := n - circle - 2; i > circle; i-- {
				result[i][circle] = start
				start++
			}
			circle++
		}
	}
	return result
}
