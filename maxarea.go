package main

import (
	"fmt"
)

func main() {
  //盛最多水的容器
	//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(MaxAreaDoublePtr(height))
}


//盛最多水的容器
//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
//在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

//MaxArea 盛最多水的容器,暴力解法
func MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	area := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			min := height[i]
			if height[j] < min {
				min = height[j]
			}
			width := j - i
			if area < min*width {
				area = min * width
			}
		}
	}
	return area
}

//双指针解法 ：https://zhuanlan.zhihu.com/p/64260453
func MaxAreaDoublePtr(height []int) int {
	if len(height) < 2 {
		return 0
	}
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {
		area := r - l
		if height[l] < height[r] {
			area *= height[l]
			l++
		} else {
			area *= height[r]
			r--
		}
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}
