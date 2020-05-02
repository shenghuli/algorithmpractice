package main

import (
	"fmt"
  "math"
)

func main (){
	//Pow(x, n)
	fmt.Println(mypow.MyPow(2, -2))
}

//50. Pow(x, n)
//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//
//示例 1:
//
//输入: 2.00000, 10
//输出: 1024.00000
//思路：https://zhuanlan.zhihu.com/p/59657875

//MyPow 原始解法
func MyPow(x float64, n int) float64 {
	//处理x特殊情况
	if x == 1 {
		return 1
	}
	if x == -1 {
		if n&1 != 0 {
			return -1
		}
		return 1
	}

	//处理n特殊情况
	if n == 0 {
		return 1
	}
	//处理最小负数
	if n == math.MinInt32 {
		if x > -1 && x < 1 {
			return math.MaxFloat64
		}
		return 0
	}

	//正常情况
	//原始解法
	//result := 1.0
	//if n > 0 {
	//	for i := 0; i < n; i++ {
	//		result *= x
	//	}
	//} else {
	//	n = -n
	//	for i := 0; i < n; i++ {
	//		result *= x
	//	}
	//	result = 1 / result
	//}

	//递归解法
	//if n > 0 {
	//	return powCore(x, n)
	//}
	//n = -n
	//return 1 / powCore(x, n)

	//迭代解法
	if n > 0 {
		return powCore2(x, n)
	}
	n = -n
	return 1 / powCore2(x, n)
}

//powCore 递归解法
func powCore(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		return powCore(x*x, n/2)
	}
	return x * powCore(x*x, n/2)
}

//powCore2 迭代解法
func powCore2(x float64, n int) float64 {
	//x^10 = x的(1010）2次方 =x^(1*2^3+0*2^2+1*2^1+0*2^0)
	//=x^(2^3)*x^(2^1)

	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x = x * x
		n = n >> 1
	}
	return result
}
