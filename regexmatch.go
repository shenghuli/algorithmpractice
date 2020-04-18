package main

//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

import (
	"fmt"
)

func main() {
	//正则表达式匹配
	//s := []string{"aa", "aa", "ab", "aab", "mississippi"}
	//p := []string{"a", "a*", ".*", "c*a*b", "mis*is*p*."}
	//s := []string{"ab", "a", ""}
	//p := []string{".*c", "ab*", ".*"}
	s := []string{""}
	p := []string{".*"}
	for i := 0; i < len(s); i++ {
		fmt.Println(IsMatch(s[i], p[i]))
	}
  
}

//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

//IsMatch 字符串是否匹配 ，剑指offer 第19道题解法
func IsMatch(s string, p string) bool {
	return matchCore(s, 0, p, 0)
}

func matchCore(s string, sstart int, p string, pstart int) bool {
	if pstart == len(p) { //正则匹配完成
		if sstart == len(s) { //字符串也匹配完成，则匹配
			return true
		} else { //字符串还有剩余，则未匹配
			return false
		}
	}
	//下一个字符为*,存在回溯的情况
	if pstart+1 < len(p) && p[pstart+1] == '*' {
		if sstart == len(s) { //string匹配结束,只可能
			return matchCore(s, sstart, p, pstart+2) //忽略本次pattern尝试往后匹配
		} else if p[pstart] == s[sstart] || p[pstart] == '.' { //当前字符匹配
			return matchCore(s, sstart+1, p, pstart+2) || //只匹配当前一个
				matchCore(s, sstart+1, p, pstart) || //尝试匹配更多
				matchCore(s, sstart, p, pstart+2) //忽略本次pattern尝试往后匹配
		} else { //未匹配，将当前这对pattern对设置为无效
			return matchCore(s, sstart, p, pstart+2)
		}
	}
	//下一个字符为非*
	if sstart == len(s) { //字符串匹配完成，还有剩余的pattern
		return false
	} else if s[sstart] == p[pstart] { //当前字符匹配，移动至下一次匹配
		return matchCore(s, sstart+1, p, pstart+1)
	} else if p[pstart] == '.' { //pattern 为.,向后走一步
		return matchCore(s, sstart+1, p, pstart+1)
	}
	return false
}

//网上解法
//https://www.imooc.com/article/281353?block_id=tuijian_wz
//IsMatchRecursive 递归
func IsMatchRecursive(s string, p string) bool {
	if len(p) == 0 { //pattern 匹配完了
		if len(s) == 0 { //字符串也匹配完了
			return true
		} else { //字符串有剩余
			return false
		}
	}
	//朴素的是否匹配
	isMatched := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return IsMatchRecursive(s, p[2:]) ||
			isMatched && IsMatchRecursive(s[1:], p)
	}
	return isMatched && IsMatchRecursive(s[1:], p[1:])
}

//IsMatchDynamicUp 动态规划自上而下实现
func IsMatchDynamicUp(s, p string) bool {
	//构建状态空间
	rows := len(s) + 1
	cols := len(p) + 1
	stageMatrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		stageMatrix[i] = make([]int, cols)
	}
	return getMatchedValue(0, 0, s, p, stageMatrix)
}

//获取状态矩阵的值，0：未初始化，-1：false,1:true
func getMatchedValue(i, j int, s, p string, stagetMatrix [][]int) bool {
	if stagetMatrix[i][j] == 1 {
		return true
	} else if stagetMatrix[i][j] == -1 {
		return false
	}

	var answer bool
	if j == len(p) {
		answer = i == len(s)
	} else {
		curMatched := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			answer = getMatchedValue(i, j+2, s, p, stagetMatrix) ||
				(curMatched && getMatchedValue(i+1, j, s, p, stagetMatrix))
		} else {
			answer = curMatched && getMatchedValue(i+1, j+1, s, p, stagetMatrix)
		}
	}

	if answer == true {
		stagetMatrix[i][j] = 1
	} else {
		stagetMatrix[i][j] = -1
	}
	return answer
}

//IsMatchDynamicUp 动态规划底向上实现
func IsMatchDynamicDown(s, p string) bool {
	//构建状态空间
	rows := len(s) + 1
	cols := len(p) + 1
	stageMatrix := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		stageMatrix[i] = make([]bool, cols)
	}
	stageMatrix[rows-1][cols-1] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			isMatched := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				stageMatrix[i][j] = stageMatrix[i][j+2] || isMatched && stageMatrix[i+1][j]
			} else {
				stageMatrix[i][j] = isMatched && stageMatrix[i+1][j+1]
			}
		}
	}
	return stageMatrix[0][0]
}

