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

func main() {
str1 := "abcabcbb"
	len1 := LengthOfLongestSubstring(str1)
	fmt.Println(len1)
	str2 := "bbbbb"
	len2 := LengthOfLongestSubstring(str2)
	fmt.Println(len2)
	str3 := "eeydgwdykpv"
	len3 := LengthOfLongestSubstring(str3)
	fmt.Println(len3)
}


//LengthOfLongestSubstring 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	strLen := len(s)
	maxSubLen := 0
	maxSubLenOfI := 0
	positionDic := map[byte]int{}
	for i:=0;i<strLen;i++{
		ch := s[i]
		if index,ok := positionDic[ch];ok{
			d := i - index
			if d > maxSubLenOfI {
				maxSubLenOfI++
			}else {
				maxSubLenOfI = d
			}
		}else{
			maxSubLenOfI++;
		}
		positionDic[ch] = i
		if maxSubLenOfI > maxSubLen {
			maxSubLen = maxSubLenOfI
		}
	}
	return maxSubLen
}



