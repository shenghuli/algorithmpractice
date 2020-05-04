package main

import (
	"fmt"
)

func main(){
  //接雨水
	fmt.Println(trap.Trap4([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

//42. 接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
//上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。
//
//示例:
//
//输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
//思路：https://zhuanlan.zhihu.com/p/79811305

//trap 接雨水，解法1，按列求解
func Trap(height []int) int {
	result := 0
	for i := 1; i < len(height)-1; i++ {

		leftheight := 0
		for left := 0; left < i; left++ {
			if height[left] > leftheight {
				leftheight = height[left]
			}
		}

		rightheight := 0
		for right := i + 1; right < len(height); right++ {
			if height[right] > rightheight {
				rightheight = height[right]
			}
		}

		minheght := leftheight
		if rightheight < minheght {
			minheght = rightheight
		}

		if minheght > height[i] {
			result += minheght - height[i]
		}
	}

	return result
}

//Trap2 动态规划解法
func Trap2(height []int) int {
	result := 0
	maxleft := make([]int, len(height))
	maxright := make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		if height[i-1] > maxleft[i-1] {
			maxleft[i] = height[i-1]
		} else {
			maxleft[i] = maxleft[i-1]
		}
		rindex := len(height) - 1 - i
		if height[rindex+1] > maxright[rindex+1] {
			maxright[rindex] = height[rindex+1]
		} else {
			maxright[rindex] = maxright[rindex+1]
		}
	}
	for i := 1; i < len(height)-1; i++ {
		minheight := maxleft[i]
		if minheight > maxright[i] {
			minheight = maxright[i]
		}
		if minheight > height[i] {
			result += minheight - height[i]
		}
	}
	return result
}

//Trap3 双指针解法
func Trap3(height []int) int {
	result := 0
	maxleft := 0
	maxright := make([]int, len(height))
	for i := len(height) - 2; i > 0; i-- {
		if height[i+1] > maxright[i+1] {
			maxright[i] = height[i+1]
		} else {
			maxright[i] = maxright[i+1]
		}
	}
	for i := 1; i < len(height)-1; i++ {
		if maxleft < height[i-1] {
			maxleft = height[i-1]
		}
		minheight := maxleft
		if minheight > maxright[i] {
			minheight = maxright[i]
		}
		if minheight > height[i] {
			result += minheight - height[i]
		}
	}
	return result
}

//Trap4 栈解法
func Trap4(height []int) int {
	result := 0
	stack := []int{}
	current := 0
	for current < len(height) {
		for len(stack) > 0 && height[current] > height[stack[len(stack)-1]] {
			//出栈
			h := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			nowTop := stack[len(stack)-1]
			minheight := height[current]
			if minheight > height[nowTop] {
				minheight = height[nowTop]
			}
			distance := current - nowTop - 1
			water := (minheight - height[h]) * distance
			result += water
		}
		stack = append(stack, current)
		current++
	}
	return result
}
