package main

import "fmt"

func main() {
	str := "-2147483647"
	fmt.Println(Myatoi(str))
}


const (
	INT_MAX = 1<<31 - 1
	INT_MIN = -(1 << 31)
)

func Myatoi(str string) int {
	//过滤无效字符
	index := 0
	//判断正负
	sign := 1
	for ; index < len(str); index++ {
		if str[index] == ' ' {
			continue
		}
		if str[index] == '-' {
			sign = -1
			index++
			break
		}
		if str[index] == '+' {
			index++
			break
		}
		//第一个是非法字符
		if str[index] < '0' || str[index] > '9' {
			return 0
		} else { //出现第一个正常数字
			break
		}
	}
	//无有效字符
	if index >= len(str) {
		return 0
	}
	retNum := 0
	for ; index < len(str); index++ {
		//正常数字
		if str[index] >= '0' && str[index] <= '9' {
			//越界处理
			bit := int(str[index] - '0')
			if sign > 0 { //正数
				if retNum > INT_MAX/10 {
					return INT_MAX
				} else if retNum == INT_MAX/10 && bit > 7 {
					return INT_MAX
				}
			} else { //负数
				if retNum < INT_MIN/10 {
					return INT_MIN
				} else if retNum == INT_MIN/10 && bit > 8 {
					return INT_MIN
				}
			}
			retNum = retNum*10 + sign*bit
		} else { //出现非正常数字，则不在匹配
			break
		}
	}
	return retNum
}
