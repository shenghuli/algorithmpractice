package main

import (
	"fmt"
)

func main(){
	//合并两个有序数组
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	Merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

//88. 合并两个有序数组
//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//
//示例:
//
//输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if m < 0 || n < 0 { //入参非法
		return
	}
	r1 := m - 1
	r2 := n - 1
	for i := m + n - 1; i >= 0; {
		//处理第二个数组
		for r2 >= 0 && (r1 < 0 || nums2[r2] >= nums1[r1]) {
			nums1[i] = nums2[r2]
			r2--
			i--
		}
		//如果第二个数组为空，则无需再排列
		if r2 < 0 {
			break
		}

		//放置第一个数组
		for r1 >= 0 && nums1[r1] >= nums2[r2] {
			nums1[i] = nums1[r1]
			r1--
			i--
		}
	}
}
