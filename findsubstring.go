package main

import (
	"fmt"
  "reflect"
)

func main(){
  //串联所有单词的子串
	fmt.Println(FindSubstring2("barfoothefoobarman", []string{"foo", "bar"}))
}

//30. 串联所有单词的子串
//给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//
//注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
//
//
//
//示例 1：
//
//输入：
//s = "barfoothefoobarman",
//words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。

//FindSubstring map解法
func FindSubstring(s string, words []string) []int {
	if len(words) < 1 || len(s) < 1 {
		return []int{}
	}
	result := []int{}
	wordLen := len(words[0])
	totalLen := len(words) * wordLen
	wordsMap := map[string]int{}
	for i := 0; i < len(words); i++ {
		k := words[i]
		if count, has := wordsMap[k]; has {
			wordsMap[k] = count + 1
		} else {
			wordsMap[k] = 1
		}
	}
	for i := 0; i <= len(s)-totalLen; i++ {
		substr := s[i : i+totalLen]
		strMap := map[string]int{}
		for j := 0; j < totalLen; j += wordLen {
			w := substr[j : j+wordLen]
			if count, has := strMap[w]; has {
				strMap[w] = count + 1
			} else {
				strMap[w] = 1
			}
		}
		if reflect.DeepEqual(wordsMap, strMap) {
			result = append(result, i)
		}
	}
	return result
}

//FindSubstring 滑动窗口解法
func FindSubstring2(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	wordMap := map[string]int{}
	for i := 0; i < len(words); i++ {
		k := words[i]
		if v, has := wordMap[k]; has {
			wordMap[k] = v + 1
		} else {
			wordMap[k] = 1
		}
	}
	wordLen := len(words[0])
	result := []int{}
	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		count := 0
		//窗口
		window := map[string]int{}
		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right = right + wordLen
			counter, has := wordMap[word]
			if !has { //在字典中不存在，重新计数
				left = right
				count = 0
				window = map[string]int{} //清空window
			} else {
				//进入滑动窗口
				count++
				if _, ok := window[word]; ok {
					window[word] += 1
					//确保新入滑动窗口的word的个数不高于words中的个数
					for counter < window[word] {
						lword := s[left : left+wordLen]
						count--
						window[lword] -= 1
						left += wordLen
					}
				} else {
					window[word] = 1
				}
			}
			if count == len(words) {
				result = append(result, left)
			}
		}
	}
	return result
}


