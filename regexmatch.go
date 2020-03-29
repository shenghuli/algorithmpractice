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
