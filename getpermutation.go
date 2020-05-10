package main

import (
	"fmt"
)

func main(){
  //第k个排列
	fmt.Println(GetPermutation(3, 2))
}

//60. 第k个排列

//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//
//按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//给定 n 和 k，返回第 k 个排列。
//
//说明：
//
//给定 n 的范围是 [1, 9]。
//给定 k 的范围是[1,  n!]。

//大体思路：决定第i个元素放置什么数，依次取决于k/（n-i-1)!的个数的结果来判断放置第几大数

//GetPermutation 第k个排列
func GetPermutation(n int, k int) string {
	chosen := make([]bool, n)
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < len(factorial); i++ {
		factorial[i] = i * factorial[i-1]
	}
	result := []byte{}
	k-- //从0开始
	for i := n - 1; i >= 0; i-- {
		index := -1
		count := 0 //第几大数
		repeatNum := k/factorial[i] + 1
		for count < repeatNum {
			index++
			if !chosen[index] {
				count++
			}
		}
		if index == -1 {
			break
		}
		chosen[index] = true
		result = append(result, byte(index+1+'0'))
		k %= factorial[i]
	}
	for i := 0; i < len(chosen); i++ {
		if !chosen[i] {
			result = append(result, byte(i+1+'0'))
		}
	}
	return string(result)
}
