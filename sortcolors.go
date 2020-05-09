package main

import (
	"fmt"
)


func main(){
	//颜色分类
	nums := []int{0, 1, 2, 1, 1, 0}
	SortColors(nums)
	fmt.Println(nums)
}


//75. 颜色分类
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//注意:
//不能使用代码库中的排序函数来解决这道题。
//
//示例:
//
//输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]

//sortColors 颜色分类，双指针解法
func SortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for i := 0; i <= right; {
		if nums[i] == 0 {
			if i == left {
				i++
			} else {
				nums[left], nums[i] = nums[i], nums[left]
			}
			left++
		} else if nums[i] == 2 && i < right {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}
}
