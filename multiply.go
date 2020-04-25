package main

import (
	"fmt"
)

func main(){
	//字符串相乘
	fmt.Println(Multiply("123", "456"))
}





//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"

//解法：https://zhuanlan.zhihu.com/p/90662602

//Multiply ，思路，将两数相乘转换为对应的位数相乘
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ret := make([]uint8, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- { //从右向左
		ni := num2[i] - '0'
		for j := len(num1) - 1; j >= 0; j-- { //从右向左
			//两位相乘，其对应的位数为i+j,i+j+1
			nj := num1[j] - '0'
			sum := ret[i+j+1] + ni*nj
			ret[i+j+1] = sum % 10
			ret[i+j] += sum / 10
		}
	}
	if ret[0] == 0 {
		ret = ret[1:]
	}
	retByte := make([]byte, len(ret))
	for i := 0; i < len(ret); i++ {
		retByte[i] = ret[i] + '0'
	}
	return string(retByte)
}
