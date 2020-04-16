package main

import (
	"fmt"
  "strings"
)

func main(){

  //判断子序列
	fmt.Println(issubsequence.IsSubsequence("abc", "axxsdsdbafdc"))

}








//392. 判断子序列
//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
//你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
//
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

//IsSubsequence 判断子序列，贪心算法
func IsSubsequence(s, t string) bool {
	for i := 0; i < len(s); i++ {
		p := strings.IndexByte(t, s[i])
		if p < 0 {
			return false
		}
		t = t[p+1:]
	}
	return true
}

//IsSubsequence 判断子序列，双指针解法
func IsSubsequenceDoubblePtr(s, t string) bool {
	if len(s) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
	}

	return false
}
