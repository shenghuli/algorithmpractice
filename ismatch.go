package main

import (
	"fmt"
)

func main(){
  fmt.Println(IsMatch("agdd","a*"))
}

//isMatchDynamicUpMa 马老师动态规划解法
func IsMatchDynamicUpMa(s, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}
	slen := len(s)
	plen := len(p)
	//dp[i,j] s的前i个元素是否被p的前j个元素成功匹配
	dp := make([][]bool, slen+1)
	for i := 0; i <= slen; i++ {
		dp[i] = make([]bool, plen+1)
	}
	dp[0][0] = true
	//初始化包含*的部分
	for i := 1; i <= plen; i++ {
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}
	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			si := s[i-1]
			pj := p[j-1]
			if si == pj || pj == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if pj == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[slen][plen]
}

//IsMatch 双索引法
func IsMatch(s, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}
	i, j := 0, 0
	istart, jsstart := -1, -1
	plen := len(p)
	for i < len(s) {
		if j < plen && (s[i] == p[j] || p[j] == '?') { //正常匹配
			i++
			j++
		} else if j < plen && p[j] == '*' { //出现*的情况
			istart = i
			jsstart = j
			j++
		} else if istart > -1 { //s匹配尽可能多的*
			i = istart + 1
			istart = i
			j = jsstart + 1
		} else { //匹配失败
			return false
		}
	}

	//p中剩余*的处理
	for j < plen && p[j] == '*' {
		j++
	}
	return j == plen
}
