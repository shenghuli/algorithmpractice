package main

import (
	"fmt"
)

func main(){
  //63. 不同路径 II
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},

		//[]int{1},
		//[]int{0},
	}
	fmt.Println(UniquePathsWithObstacles(grid))
}


//63. 不同路径 II
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//说明：m 和 n 的值均不超过 100。
//
//示例 1:
//
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右

//UniquePathsWithObstacles 不同路径 II
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 || len(obstacleGrid[0]) < 1 {
		return 0
	}
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	flag := false
	for i := 0; i < cols; i++ {
		if obstacleGrid[0][i] == 1 {
			flag = true
		}
		if flag {
			obstacleGrid[0][i] = -1
		} else {
			obstacleGrid[0][i] = 1
		}
	}
	flag = false
	if obstacleGrid[0][0] == -1 {
		flag = true
	}
	for i := 1; i < rows; i++ {
		if obstacleGrid[i][0] == 1 {
			flag = true
		}
		if flag {
			obstacleGrid[i][0] = -1
		} else {
			obstacleGrid[i][0] = 1
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			} else {
				top := obstacleGrid[i-1][j]
				left := obstacleGrid[i][j-1]
				if top == -1 && left == -1 {
					obstacleGrid[i][j] = -1
				}
				if top == -1 {
					top = 0
				}
				if left == -1 {
					left = 0
				}
				obstacleGrid[i][j] = top + left
			}
		}
	}
	if obstacleGrid[rows-1][cols-1] == -1 {
		return 0
	}
	return obstacleGrid[rows-1][cols-1]
}
