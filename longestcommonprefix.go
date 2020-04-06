package main

import (
	"fmt"
)

func main() {

	//最长公共前缀
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	//strs := []string{"b", ""}
	strs := []string{"aa", "a"}
	fmt.Println(LongestCommonPrefix(strs))

}


//最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 { //空数组
		return ""
	}
	if len(strs) == 1 { //只有一个
		return strs[0]
	}
	i := 0
	for ; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return strs[0][0:i]
			}
			if strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0][0:i]
}


//LongestCommonPrefix2 两两归并方法
func LongestCommonPrefix2(strs []string) string {
	if len(strs) < 1 { //空数组
		return ""
	}
	if len(strs) == 1 { //只有一个
		return strs[0]
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result = twoLonggetCommonPrefix(result, strs[i])
	}
	return result
}

//twoLonggetCommonPrefix
func twoLonggetCommonPrefix(str1, str2 string) string {
	minLen := len(str1)
	if minLen > len(str2) {
		minLen = len(str2)
	}
	index := -1
	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			break
		}
		index++
	}
	if index == -1 {
		return ""
	}
	return str1[:index+1]
}
