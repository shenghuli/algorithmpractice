package main

import "fmt"

func main(){
  //合并K个有序链表
	l1 := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  11,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 10,
		},
	}
	list := []*ListNode{nil, l1, nil, l2}
	mergeList := MergeKLists2(list)
	ToString(mergeList)

}


//ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoList 循环解法
func MergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	pHead := &ListNode{Val: 0}
	tmpNodeList := pHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmpNodeList.Next = l1
			l1 = l1.Next
		} else {
			tmpNodeList.Next = l2
			l2 = l2.Next
		}
		tmpNodeList = tmpNodeList.Next
	}
	if l1 == nil {
		tmpNodeList.Next = l2
	} else {
		tmpNodeList.Next = l1
	}
	return pHead.Next
}

//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
//示例:
//
//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6

//MergeKLists 合并K个排序链表,解法一两两合并，参见mergetwolist
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 { //列表为空
		return nil
	}
	if len(lists) == 1 { //长度为1
		return lists[0]
	}
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = MergeTwoList(result, lists[i])
	}
	return result
}

//MergeKLists 合并K个排序链表,每次获取最小的值
func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 { //列表为空
		return nil
	}
	if len(lists) == 1 { //长度为1
		return lists[0]
	}
	pHead := &ListNode{}
	pTmpList := pHead
	INT_MAX := int(^uint(0) >> 1)
	for {
		//int_max
		min := INT_MAX
		minIndex := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if lists[i].Val < min {
					min = lists[i].Val
					minIndex = i
				}
			}
		}
		if minIndex == -1 {
			break
		}
		pTmpList.Next = lists[minIndex]
		pTmpList = pTmpList.Next
		lists[minIndex] = lists[minIndex].Next
	}
	return pHead.Next
}

//ToString 打印单链表
func ToString(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}



