/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	var retArr []int

	retArr = twoSum(nums,target)
	fmt.Println(retArr)
}

//暴力实现
func twoSum(nums []int, target int) []int {
    count := len(nums)
    var retArr []int
    for i := 0; i< count ; i++ {
    	for j := i+1; j < count; j++ {
    		if nums[i] + nums[j] == target {
    			retArr = append(retArr,i,j)
    			return retArr
    		}
    	}
    }

    return []int{}
}

//map实现
func twoSum(nums []int, target int) []int {
    var tmpDic = map[int]int{}
	//构建字典
	for i := 0; i < len(nums); i++ {
		k := nums[i]
		v := i
		tmpDic[k] = v
	}
	//map查询
	for i := 0; i < len(nums); i++ {
		tk := target - nums[i]
		if ti, ok := tmpDic[tk]; ok {
			return []int{i, ti}
		}
	}
	return []int{}
}
