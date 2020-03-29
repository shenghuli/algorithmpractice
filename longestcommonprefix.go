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
