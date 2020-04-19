package main

import (
	"fmt"
)

func main(){
  //实现 strStr()
	fmt.Println(strstr.StrStr("aaa", "aaaa"))
}


//28. 实现 strStr()
//实现 strStr() 函数。
//
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1:
//
//输入: haystack = "hello", needle = "ll"
//输出: 2
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		tmpI := i
		isContain := true
		for j := 0; j < len(needle); j++ {
			if tmpI >= len(haystack) {
				return -1
			}
			if haystack[tmpI] != needle[j] {
				isContain = false
				break
			}
			tmpI++
		}
		if isContain {
			return i
		}
	}
	return -1
}
