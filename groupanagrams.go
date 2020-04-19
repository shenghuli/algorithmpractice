package main

import (
	"fmt"
  "sort"
)

func main(){
  //字母异位词分组
	fmt.Println(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}


//49. 字母异位词分组
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。

type SortString []byte

func (str SortString) Len() int {
	return len(str)
}
func (str SortString) Less(i, j int) bool {
	return str[i] < str[j]
}
func (str SortString) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}

func GroupAnagrams(strs []string) [][]string {
	sortStrDic := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		sortstr := SortString(str)
		sort.Sort(sortstr)
		key := string(sortstr)
		if list, has := sortStrDic[key]; !has {
			if list == nil {
				sortStrDic[key] = []string{str}
			}
		} else {
			sortStrDic[key] = append(sortStrDic[key], str)
		}
	}
	retList := [][]string{}
	for _, list := range sortStrDic {
		retList = append(retList, list)
	}

	return retList
}
