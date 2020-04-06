package main

import (
	"fmt"
)

func main(){
	//求链表的中间节点
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	tmpNode := FindMidlle(l1)
	fmt.Println(tmpNode.Val)
}




//ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//选自剑指offer第22题相关题目，题目：
//求链表的中间节点。如果链表中的节点总数为奇数，则返回中间节点；
//如果节点总数为偶数，则返回中间节点中的任意一个

//FindMidlle 求链表的中间节点
func FindMidlle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pFast := head
	pSlow := head
	if pFast.Next != nil {
		pFast = pFast.Next.Next
		pSlow = pSlow.Next
	}
	return pSlow
}
