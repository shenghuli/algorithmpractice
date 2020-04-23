package main

import (
  "fmt"
)

func main(){
  //两数相除
	fmt.Println(DivideBit(10, 3))
}

//29. 两数相除
//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数 dividend 除以除数 divisor 得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
//示例 1:
//
//输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3

const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = -INT_MAX - 1

//Divide 两数相除，相减法
func Divide(dividend int, divisor int) int {
	if divisor == 0 { //异常
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	//处理越界的情况
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	absDividend := abs(dividend)
	absDivisor := abs(divisor)
	//模拟普通的除法
	shang := 0
	for absDividend >= absDivisor {
		absDividend = absDividend - absDivisor
		shang++
	}
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		return -shang
	}
	return shang
}

//DivideBit 两数相除，位运算
func DivideBit(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}
	absDividend := abs(dividend)
	absDivisor := abs(divisor)
	//dividend = divisor*an^n+divisor*an-1*2^(n-1)+...+a0*2^0 + ys
	shang := 0
	for absDividend >= absDivisor {
		//寻找当前最大基的系数
		tmp := absDivisor //2^0
		tmpShang := 1     //最大基对应的商
		for absDividend > tmp<<1 {
			tmp = tmp << 1
			tmpShang = tmpShang << 1
		}
		absDividend -= tmp
		shang += tmpShang
	}
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		return -shang
	}
	return shang
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
