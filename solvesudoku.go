package main

import (
	"fmt"
)

func main(){
  //解数独
	sudoku := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	SolveSudoku(sudoku)
	fmt.Println(sudoku)
}


//37. 解数独
//编写一个程序，通过已填充的空格来解决数独问题。
//
//一个数独的解法需遵循如下规则：
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
//空白格用 '.' 表示。
//
//
//
//一个数独。
//思路：https://zhuanlan.zhihu.com/p/128066021

//SolveSudoku 大体思路，构造行、列、子块各个数字对应个数的数组
//然后逐步回溯
func SolveSudoku(board [][]byte) {
	rows := make([][]int, 9)
	columns := make([][]int, 9)
	boxes := make([][]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]int, 10) //用0-9当索引
		columns[i] = make([]int, 10)
		boxes[i] = make([]int, 10)
	}
	//初始化行、列、子块
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			char := int(board[i][j] - '0')
			rows[i][char] = 1
			columns[j][char] = 1
			boxindex := i/3*3 + j/3
			boxes[boxindex][char] = 1
		}
	}
	isSolved := false
	backtrack(0, 0, &rows, &columns, &boxes, &board, &isSolved)
}

//backtrack 回溯法
func backtrack(row, column int, rows, columns, boxes *[][]int, board *[][]byte, isSolved *bool) {
	if (*board)[row][column] == '.' {
		boxindex := row/3*3 + column/3
		for i := 1; i < 10; i++ {
			if (*rows)[row][i]+(*columns)[column][i]+(*boxes)[boxindex][i] == 0 {
				(*rows)[row][i]++
				(*columns)[column][i]++
				(*boxes)[boxindex][i]++
				(*board)[row][column] = byte(i) + '0'
				if row == 8 && column == 8 {
					*isSolved = true
				} else if column == 8 {
					backtrack(row+1, 0, rows, columns, boxes, board, isSolved)
				} else {
					backtrack(row, column+1, rows, columns, boxes, board, isSolved)
				}
				//回溯
				if *isSolved == false {
					(*rows)[row][i]--
					(*columns)[column][i]--
					(*boxes)[boxindex][i]--
					(*board)[row][column] = '.'
				}
			}
		}
	} else {
		if row == 8 && column == 8 { //到达了右下角
			*isSolved = true
			return
		}
		if column == 8 { //到达了右边
			backtrack(row+1, 0, rows, columns, boxes, board, isSolved)
		} else {
			backtrack(row, column+1, rows, columns, boxes, board, isSolved)
		}
	}
}

func printByteArr(arr [][]byte) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(string(arr[i][j]), " ")
		}
		fmt.Println()
	}
}
