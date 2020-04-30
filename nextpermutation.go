package main

import (
	"fmt"
)

func main(){
	//下一个排列
	nums := []int{1, 5, 1}
	NexPermutation(nums)
	fmt.Println(nums)
}

//31. 下一个排列
//实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须原地修改，只允许使用额外常数空间。
//
//以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

// 大体思路,从右向左找到第一个比右边小的数然后和将它插入到右边序列中合适的位置
// 最后将右边的序列逆序输出，即可得到,参考解法：https://zhuanlan.zhihu.com/p/45007701
//NexPermutation 下一个排列
func NexPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	//寻找第一个比右边小的数
	index := len(nums) - 2
	for ; index >= 0; index-- {
		if nums[index] < nums[index+1] {
			break
		}
	}
	//寻找右侧比index小的数，其前一个就是要交换的位置
	if index >= 0 {
		swapIndex := index + 1
		for swapIndex < len(nums) {
			if nums[swapIndex] <= nums[index] {
				break
			}
			swapIndex++
		}
		//交换
		tmp := nums[index]
		nums[index] = nums[swapIndex-1]
		nums[swapIndex-1] = tmp
	}
	//反转
	left := index + 1
	right := len(nums) - 1
	for left < right {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
		left++
		right--
	}
}
