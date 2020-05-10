package mergerange

import (
	"fmt"
  "sort"
)

func main(){
  //合并区间
	intRange := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	fmt.Println(Merge(intRange))
}

//56. 合并区间
//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

type RangeSort [][]int

func (r RangeSort) Len() int {
	return len(r)
}

func (r RangeSort) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RangeSort) Less(i, j int) bool {
	return r[i][0] < r[j][0]
}

//Merge 给出一个区间的集合，请合并所有重叠的区间。
func Merge(intervals [][]int) [][]int {
	result := [][]int{}
	sort.Sort(RangeSort(intervals))
	for i := 0; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]
		for i+1 < len(intervals) && intervals[i+1][0] <= right {
			if intervals[i+1][1] > right {
				right = intervals[i+1][1]
			}
			i++
		}
		result = append(result, []int{left, right})
	}
	return result
}
