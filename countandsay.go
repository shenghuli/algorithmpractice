package main

import (
	"fmt"
)

func main(){

	//搜索插入位置
	fmt.Println(SearchInsert([]int{1, 3}, 3))

}


//38. 外观数列
//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
//注意：整数序列中的每一项将表示为一个字符串。
//
//
//
//示例 1:
//
//输入: 1
//输出: "1"
//解释：这是一个基本样例。

func CountAndSay(n int) string {
	if n < 1 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	str := CountAndSay(n - 1)
	retStr := ""
	count := 1
	prev := str[0]
	for i := 1; i < len(str); i++ {
		if str[i] == prev {
			count++
		} else {
			retStr += fmt.Sprintf("%d", count) + string(prev)
			prev = str[i]
			count = 1
		}
	}
	retStr += fmt.Sprintf("%d", count) + string(prev)
	return retStr
}
