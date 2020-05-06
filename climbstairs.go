package main

import (
	"fmt"
)

func main(){
  //爬楼梯
	fmt.Println(ClimbStairs(3))
}

//70. 爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶

//解法：剑指offer 第10题

//ClimbStairs 爬楼梯,斐波那契数列
func ClimbStairs(n int) int {
	if n <= 1 {
		return n
	}
	fminus2 := 1
	fminus1 := 1
	for i := 2; i <= n; i++ {
		f := fminus1 + fminus2
		fminus2 = fminus1
		fminus1 = f
	}

	return fminus1
}
