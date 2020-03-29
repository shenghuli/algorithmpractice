package main

import (
	"fmt"
)

func main() {
	//罗马数字转整数
	//str := "LVIII"
	s := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for i := 0; i < len(s); i++ {
		fmt.Println(romantoint.RomanToInt(s[i]))
	}
	//str := "IV"
	//fmt.Println(RomanToInt(str))
}


var charmap map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	retInt := 0
	for i := 0; i < len(s); i++ {
		v := charmap[s[i]]
		if i+1 < len(s) {
			vnext := charmap[s[i+1]]
			if vnext > v {
				retInt += (vnext - v)
				i++
				continue
			}
		}
		retInt += v
	}
	return retInt
}
