package main

import "fmt"

func main() {
  //合并两个有序链表
	l1 := &mergetwolist.ListNode{
		Val: 1,
		Next: &mergetwolist.ListNode{
			Val: 2,
			Next: &mergetwolist.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &mergetwolist.ListNode{
		Val: 1,
		Next: &mergetwolist.ListNode{
			Val: 3,
			Next: &mergetwolist.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	mergeList := MergeTwoList(l1, l2)
	ToString(mergeList)
}

//21. 合并两个有序链表
//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

//ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoList 剑指Offer第25题解法，递归解法
func MergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var mergeList *ListNode
	if l1.Val <= l2.Val {
		mergeList = l1
		mergeList.Next = MergeTwoList(l1.Next, l2)
	} else {
		mergeList = l2
		mergeList.Next = MergeTwoList(l1, l2.Next)
	}
	return mergeList
}

//MergeTwoList 循环解法
func MergeTwoList2(l1 *ListNode, l2 *ListNode) *ListNode {
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

//ToString 打印单链表
func ToString(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}

  
  
  
