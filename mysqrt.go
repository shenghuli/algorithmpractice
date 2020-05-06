package main

import (
	"fmt"
)

func main(){
  //x的平方根
	fmt.Println(MySqrt2(8))
}


//x 的平方根
//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2

//牛顿法：https://zhuanlan.zhihu.com/p/37588590
//牛顿法：https://zhuanlan.zhihu.com/p/74478033

//MySqrt x的平方根，牛顿法
//https://zhuanlan.zhihu.com/p/67293542
func MySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

//MySqrt x的平方根，二分法
func MySqrt2(x int) int {
	if x <= 1 {
		return x
	}
	left := 1
	right := (x + 1) / 2 //几何平均数不大于算术平均数
	for left <= right {  //最终right会走到left左边
		mid := (left + right) / 2
		tmp := x / mid
		if mid == tmp {
			return mid
		}
		if mid < tmp {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right //返回较小的
}
