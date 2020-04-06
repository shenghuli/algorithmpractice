package main

import (
	"fmt"
)


func main() {
  //有效的括号
	strs := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	//strs := []string{"()"}
	for i := 0; i < len(strs); i++ {
		fmt.Println(IsValid(strs[i]))
	}
}










//20. 有效的括号

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。

// IsValid 本质上考察的是栈
func IsValid(s string) bool {
	stack := []byte{}
	dic := map[byte]int{
		'(': 1,
		'{': 2,
		'[': 3,
		')': -1,
		'}': -2,
		']': -3,
	}
	//非偶数个字符，必定为false
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if v, ok := dic[s[i]]; ok {
			if v > 0 {
				stack = append(stack, s[i])
			} else {
				if len(stack) < 1 {
					return false
				}
				ch := stack[len(stack)-1]
				if v+dic[ch] != 0 {
					return false
				}
				//出栈
				stack = stack[:len(stack)-1]
			}
		} else {
			//非法输入
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
