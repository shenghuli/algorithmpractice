package main

import (
	"fmt"
)


const (
	INT_MAX = 2147483647
	INT_MIN = -2147483648
)

//Reverse 整数翻转 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
func Reverse(x int) int {
	var rev int = 0
	for x != 0 {
		bit := x % 10
		x = x / 10
		if rev > INT_MAX/10 || (rev == INT_MAX/10 && bit > 7) {
			return 0
		}
		if rev < INT_MIN/10 || (rev == INT_MIN/10 && bit < -8) {
			return 0
		}
		rev = rev*10 + bit
	}
	return rev
}



func main() {
  	num := 153423646
	  rnum := reverse.Reverse(num)
	  fmt.Println(rnum)
}
