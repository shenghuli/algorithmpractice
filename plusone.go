package main

import (
	"fmt"
)
 
func main {
	//加一
	fmt.Println(PlusOne([]int{9, 9, 9}))
} 

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。

func PlusOne(digits []int) []int {
	nextDigit := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if nextDigit == 0 {
			break
		}
		current := nextDigit + digits[i]
		digits[i] = current % 10
		if current == 10 {
			nextDigit = 1
		} else {
			nextDigit = 0
		}
	}
	if nextDigit == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
