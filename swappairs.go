package main

import (
	"fmt"
)

func main(){
  //两两交换链表
	root := swappairs.ListNode{
		Val: 1,
		Next: &swappairs.ListNode{
			Val: 2,
			Next: &swappairs.ListNode{
				Val: 3,
				Next: &swappairs.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(SwapPairs(&root))
}










//24. 两两交换链表中的节点
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	retHead := &ListNode{}
	retHead.Next = head
	tmpList := retHead
	for tmpList.Next != nil && tmpList.Next.Next != nil {
		//两两交换
		tmpNode := tmpList.Next
		tmpList.Next = tmpNode.Next
		tmpNode.Next = tmpList.Next.Next
		tmpList.Next.Next = tmpNode
		tmpList = tmpList.Next.Next
	}
	return retHead.Next
}
