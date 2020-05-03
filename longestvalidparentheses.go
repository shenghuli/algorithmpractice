package main

import (
	"fmt"
)
func main(){
	//最长有效括号
	fmt.Println(LongestValidParentheses2("(())()"))
}

//32. 最长有效括号
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
//示例 1:
//
//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//参考：https://zhuanlan.zhihu.com/p/64613851

//LongestValidParentheses 最长有效括号,栈解法
func LongestValidParentheses(s string) int {
	if len(s) <= 0 {
		return 0
	}
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				currentLen := i - stack[len(stack)-1]
				if result < currentLen {
					result = currentLen
				}
			}
		}
	}
	return result
}

//LongestValidParentheses 最长有效括号,动态规划解法
func LongestValidParentheses2(s string) int {
	if len(s) <= 0 {
		return 0
	}
	dp := make([]int, len(s))
	result := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 > 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if result < dp[i] {
				result = dp[i]
			}
		}
	}
	return result
}
