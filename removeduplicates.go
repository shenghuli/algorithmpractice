package main

import (
	"fmt"
)

func main() {
	//删除排序数组中的重复项
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeduplicates.RemoveDuplicates(nums))
	fmt.Println(nums)
}


//RemoveDuplicates 删除有序数组中的重复项，返回子数组的项
func RemoveDuplicates(nums []int) int {
	//双指针解法
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
