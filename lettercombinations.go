package main

import (
	"fmt"
)


func main() {
	//电话号码的字母组合
	//str := "23"
	//fmt.Println(LetterCombinations(str))
}


var phoneNumDic = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

//LetterCombinations 电话号码的字母组合
func LetterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	strs := []string{""}
	for i := 0; i < len(digits); i++ {
		tmp := phoneNumDic[digits[i]]
		tstrs := []string{}
		for j := 0; j < len(strs); j++ {
			for k := 0; k < len(tmp); k++ {
				tstrs = append(tstrs, strs[j]+string(tmp[k]))
			}
		}
		strs = tstrs
	}
	return strs
}
