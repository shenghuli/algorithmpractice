package main

import (
	"fmt"
)

func main(){
  //二进制求和
	fmt.Println(AddBinary("1010", "1011"))
}


//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
//输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"

func AddBinary(a string, b string) string {
	result := []byte{}
	i := len(a) - 1
	j := len(b) - 1
	carry := byte(0)
	for i >= 0 && j >= 0 {
		tmp := a[i] - '0' + b[j] - '0' + carry
		result = append(result, tmp%2+'0')
		carry = tmp / 2
		i--
		j--
	}

	for i >= 0 {
		tmp := a[i] - '0' + carry
		result = append(result, tmp%2+'0')
		carry = tmp / 2
		i--
	}

	for j >= 0 {
		tmp := b[j] - '0' + carry
		result = append(result, tmp%2+'0')
		carry = tmp / 2
		j--
	}
	if carry > 0 {
		result = append(result, carry+'0')
	}

	res := make([]byte, len(result))
	m := 0
	for n := len(result) - 1; n >= 0; n-- {
		res[m] = result[n]
		m++
	}

	return string(res)
}
