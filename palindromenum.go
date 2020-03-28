package main

import (
	"fmt"
）

func main() {
	x := 123321
	fmt.Println(palindromenum.IsPalindromeOptimize(x))
}

//回文数 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

//IsPalindrome 按字符的形式一位一位比对
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//计算总共有多少位
	tmp := x
	digt := 1
	size := 1
	for tmp >= 10 {
		digt++
		tmp = tmp / 10
		size *= 10
	}

	for i := 0; i < digt; i += 2 {
		//一位的数字返回true
		if digt-i == 1 {
			return true
		}
		left := x / size
		x = x % size
		size = size / 100

		right := x % 10
		x = x / 10
		if left == right {
			continue
		} else {
			return false
		}
	}
	return true
}

//IsPalindromeOptimize 优化算法O(n/2)
func IsPalindromeOptimize(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	tmpNum := 0
	for x > tmpNum {
		tmpNum = tmpNum*10 + x%10
		x = x / 10
	}
	return tmpNum == x || x == tmpNum/10
}
