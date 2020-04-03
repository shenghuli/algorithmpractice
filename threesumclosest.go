package main

import (
	"fmt"
)

func main() {
  	//最接近的三数之和
	nums := []int{-1, 2, 1, -4}
	nums := []int{-3, -2, -5, 3, -4}
	target := -1
	fmt.Println(ThreeSumClosest(nums, target))
}




const MaxInt = int(^uint(0) >> 1)

//ThreeSumClosest 三数之和,暴力解法
func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	//对原数组排序
	closestNum := MaxInt
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				//求绝对值，确定距离
				absSum := sum - target
				if absSum == 0 {
					return sum
				}
				if absSum < 0 {
					absSum = 0 - absSum
				}
				//求绝对值，确定距离
				var absClosestNum int
				if closestNum == MaxInt && target < 0 { //求距离越界
					absClosestNum = MaxInt
				} else {
					absClosestNum = closestNum - target
				}
				if absClosestNum < 0 {
					absClosestNum = 0 - absClosestNum
				}
				fmt.Println("absSum=", absSum, "absClosestNum", absClosestNum)
				if absSum < absClosestNum {
					closestNum = sum
				}
			}
		}

	}
	return closestNum
}

//ThreeSumClosest 三数之和,暴力解法2
func ThreeSumClosest1(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	//对原数组排序

	distance := float64(MaxInt)
	retInt := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum-target == 0 {
					return sum
				}
				if math.Abs(float64(sum-target)) < distance {
					distance = math.Abs(float64(sum - target))
					retInt = sum
				}
			}
		}
	}
	return retInt
}


