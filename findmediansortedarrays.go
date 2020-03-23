

package main

import (
	"fmt"
 )
 
 func main() {
  num1 := []int{100001}
	num2 := []int{100000}
	resNum := findmediansortedarrays.FindMedianSortedArrays(num1, num2)
	fmt.Println(resNum)
 }



//FindMedianSortedArrays 寻找两个有序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	l := (m + n + 1) / 2.0
	r := (m + n + 2) / 2.0
	k1 := getKthMax(nums1, 0, nums2, 0, l)
	k2 := getKthMax(nums1, 0, nums2, 0, r)
	return (k1 + k2) / 2
}

//getKthMax 从0开始，第K大的数字
func getKthMax(nums1 []int, start1 int, nums2 []int, start2 int, k int) float64 {
	len1, len2 := len(nums1), len(nums2)
	//不越界
	if start1 > len1-1 {
		return float64(nums2[start2+k-1])
	}
	if start2 > len2-1 {
		return float64(nums1[start1+k-1])
	}
	//只找一个元素
	if k == 1 {
		return math.Min(float64(nums1[start1]), float64(nums2[start2]))
	}

	nums1Mid, nums2Mid := math.MaxFloat64, math.MaxFloat64
	if start1+k/2-1 < len1 {
		nums1Mid = float64(nums1[start1+k/2-1])
	}
	if start2+k/2-1 < len2 {
		nums2Mid = float64(nums2[start2+k/2-1])
	}
	if nums1Mid < nums2Mid {
		return getKthMax(nums1, start1+k/2, nums2, start2, k-k/2)
	}
	return getKthMax(nums1, start1, nums2, start2+k/2, k-k/2)
}
