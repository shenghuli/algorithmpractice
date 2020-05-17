package main

import (
	"fmt"
)

func main(){
	list := []int{3, 1, 1, 4, 2, 6}
	qsort.Qsort(list)
	fmt.Println(list)
}


//快排

func Qsort(list []int) {
	if len(list) < 2 {
		return
	}
	qsortCore(list, 0, len(list)-1)
}

func qsortCore(list []int, left, right int) {
	temp := list[left] //分割值
	p := left          //待存放要交换元素的位置
	i, j := left, right

	for i <= j {
		for j >= p && list[j] >= temp {
			j--
		}

		if j >= p { //中间有元素小于分割值
			list[p] = list[j]
			p = j
		}

		for i <= p && list[i] <= temp {
			i++
		}
		if i <= p {
			list[p] = list[i]
			p = i
		}

		list[p] = temp
		if p-left > 1 { //分割后左边子区间大于2
			qsortCore(list, left, p-1)
		}
		if right-p > 1 { //分割后右边子区间大于2
			qsortCore(list, p+1, right)
		}
	}

}
