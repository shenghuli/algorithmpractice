package main

import (
	"fmt"
)

func main(){

  //括号生成
	fmt.Println(GenerateParenthesis(3))

}


//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例：
//
//输入：n = 3
//输出：[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]

//GenerateParenthesis 括号生成
func GenerateParenthesis(n int) []string {
	retList := []string{}
	makeGenerateParenthesis(&retList, "", 0, 0, n)
	return retList
}

//makeGenerateParenthesis 生成括号
func makeGenerateParenthesis(list *[]string, temp string, open, close, n int) {
	if open == n && close == n {
		*list = append(*list, temp)
		return
	}
	if open < n {
		makeGenerateParenthesis(list, temp+"(", open+1, close, n)
	}
	if close < open {
		makeGenerateParenthesis(list, temp+")", open, close+1, n)
	}
}
