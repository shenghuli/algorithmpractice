package main

import (
	"fmt"
)

func main(){
  //加油站
	fmt.Println(CanCompleteCircuitSum([]int{3, 3, 4}, []int{3, 4, 4}))
}

//134. 加油站
//在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
//
//说明:
//
//如果题目有解，该答案即为唯一答案。
//输入数组均为非空数组，且长度相同。
//输入数组中的元素均为非负数。
//示例 1:
//
//输入:
//gas  = [1,2,3,4,5]
//cost = [3,4,5,1,2]
//
//输出: 3
//
//解释:
//从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
//开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
//开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
//开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
//开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
//开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
//因此，3 可为起始索引。

//CanCompleteCircuit 加油站 仿真解法
func CanCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if cost[i] > gas[i] {
			continue
		}
		if isCan(gas, cost, i) {
			return i
		}
	}
	return -1
}

//isCan,以start为起点，是否能够跑完
func isCan(gas, cost []int, start int) bool {
	remain := 0
	for i := 0; i < len(gas); i++ {
		j := (start + i) % len(gas)
		remain += gas[j] - cost[j]
		if remain < 0 {
			return false
		}
	}
	return true
}

//CanCompleteCircuitSum 联系求和法
func CanCompleteCircuitSum(gas []int, cost []int) int {
	total := 0
	sum := 0
	index := 0
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff
		sum += diff
		if sum < 0 { //连续和为负数，只能出现在下一个加油站
			index = i + 1
			sum = 0
		}
	}
	if total >= 0 {
		return index
	}
	return -1
}
