package main

import (
	"fmt"
)


func main() {

  str := "PAYPALISHIRING"
	fmt.Println(ConvertOptimize(str, 3))

}

//Convert Z字变换 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

//Convert ,解法一暴力解法，寻找到周期2*numRows,然后每行输出符合的数据
func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	size := len(s)
	cycle := 2*numRows - 2 //周期
	str := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < size; j++ {
			ys := j % cycle
			if ys == i || cycle-ys == i {
				str += s[j : j+1]
			}
		}
	}
	return str
}

//ConvertOptimize 优化Convert算法，每一个周期内，首位行只打印一个元素，其他行打印两个
func ConvertOptimize(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	cycle := 2*numRows - 2
	rstr := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += cycle {
			rstr += s[j+i : j+i+1]
			//过滤掉首位行
			if i == 0 || i == numRows-1 {
				continue
			}
			//中间行的打印
			if j+cycle-i < len(s) {
				rstr += s[j+cycle-i : j+cycle-i+1]
			}
		}
	}
	return rstr
}

//ConvertByRow 按行转换
func ConvertByRow(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	minRows := numRows
	if numRows > len(s) {
		minRows = len(s)
	}
	buffList := make([]string, minRows)
	cycle := 2*numRows - 2
	for i := 0; i < len(s); i++ {
		ys := i % cycle
		if ys < numRows {
			buffList[ys] += s[i : i+1]
			continue
		}
		lineno := numRows - (ys % numRows) - 2
		buffList[lineno] += s[i : i+1]
	}
	rstr := ""
	for i := 0; i < minRows; i++ {
		rstr += buffList[i]
	}
	return rstr
}
