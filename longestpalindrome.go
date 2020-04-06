/*
题目：
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

剑指Offer中原题：
请从字符串中找出一个最长不包含重复字符的子字符串，计算该最长子字符串的长度。假设字符串中包含'a'~'z'的字符，例如，在字符串"arabcacfr"中，最长的不含重复字符的子字符串是"acfr",长度为4
*/


package main

import "fmt"

func main() {
	//最长回文子串
	str := "abcdcbafabcdck"
	lstr := LongestPalindrome(str)
	fmt.Println(lstr)
	lstr := LongestPalindromeManacher(str)
	fmt.Println(lstr)
}


//解法1，优化版暴力解法,复杂度（n^3)
func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	for size := len(s); size > 0; size-- {
		//寻找长度为size的回文
		for low, high := 0, size-1; high < len(s); low, high = low+1, high+1 {
			if isPalingdrome(s, low, high) {
				return s[low : high+1]
			}
		}
	}
	return s[0:1]
}

//解法2 LongestPalindromeCenter 中心扩散解法，复杂度(0^2)
func LongestPalindromeCenter(s string) string {
	if len(s) <= 1 {
		return s
	}
	var maxLen int = 0
	var res string = ""
	for i := 0; i < len(s)-1; i++ {
		tMaxLen, tres := expandCheckPalindrom(s, i, i, maxLen)
		if tMaxLen > maxLen {
			maxLen = tMaxLen
			res = tres
		}
		tMaxLen, tres = expandCheckPalindrom(s, i, i+1, maxLen)
		if tMaxLen > maxLen {
			maxLen = tMaxLen
			res = tres
		}
	}
	return res
}

//解法3，Manacher算法，复杂度O(n) https://www.jianshu.com/p/116aa58b7d81
var specialChar = "#"

func LongestPalindromeManacher(s string) string {
	if len(s) <= 1 {
		return s
	}
	//构建插入特殊字符的字符串
	mstr := ""
	for i := 0; i < len(s); i++ {
		mstr += specialChar + s[i:i+1]
	}
	mstr += specialChar

	//初始变量
	//C：对右回文右边界的对称中心
	//R：最右回文右边界
	//Radius：回文半径数组
	C := -1
	R := -1
	Radius := make([]int, len(mstr))
	//最大半径对应的C
	maxC := 0
	maxR := 0
	for i := 0; i < len(Radius); i++ {
		//-------- 计算Radius[i] begin
		if R > i { //case 2，
			//对称点 2*i-1
			if Radius[2*C-i] < R-i+1 { //cL > pL
				Radius[i] = Radius[2*C-i]
			} else { //cL < pL,cL=pL 起始值
				Radius[i] = R - i + 1
			}
		} else { //case 1,默认为1，起始值
			Radius[i] = 1
		}
		//以上为起始值的case，下面计算至终止态
		for i+Radius[i] < len(mstr) && i-Radius[i] > -1 {
			if mstr[i-Radius[i]] == mstr[i+Radius[i]] {
				Radius[i]++
			} else {
				break
			}
		}
		//-------- 计算Radius[i] end
		//更新R
		if R < Radius[i]+i {
			R = Radius[i] + i - 1
			//更新C
			C = i
		}
		//更新maxC
		if maxR < Radius[i] {
			maxR = Radius[i]
			maxC = i
		}
	}
	//输出起始值
	start := (maxC - Radius[maxC] + 1) / 2
	end := (maxC + Radius[maxC] - 1) / 2
	return s[start:end]
}

//expandCheckPalindrom 从中心向两遍扩展，检查是否是回文
func expandCheckPalindrom(s string, low, high, maxLen int) (retMaxLen int, res string) {
	retMaxLen = maxLen
	res = ""
	for low >= 0 && high < len(s) {
		if s[low] == s[high] {
			if high-low+1 > maxLen {
				retMaxLen = high - low + 1
				res = s[low : high+1]
			}
			low--
			high++
		} else {
			//不是回文，退出
			return retMaxLen, res
		}
	}
	return retMaxLen, res
}

//isPalingdrome 判断是否是回文
func isPalingdrome(s string, low, high int) bool {
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			return false
		}
	}
	return true
}


//状态转移数组，dp[l,r]表示s[l,r]是否构成了回文字符串
//LongestPalindrome3 最长的回文子串，解法4，动态规划
func LongestPalindrome3(s string) string {
	if len(s) <= 1 {
		return s
	}
	longestPalindromeLen := 1
	longestPalindromeStr := s[:1]
	//初始化状态转移表
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if (s[r] == s[l]) && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
				if r-l+1 > longestPalindromeLen {
					longestPalindromeLen = r - l + 1
					longestPalindromeStr = s[l : r+1]
				}
			}
		}
	}

	return longestPalindromeStr
}
