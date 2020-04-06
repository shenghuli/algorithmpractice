package main

import (
	"fmt"
)


func main(){
  //链表中倒数第k个节点
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	ToString(FindKthToTail(l1, 1))
}


//ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//FindKthToTail 链表中的倒数第K个节点
func FindKthToTail(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	pFirst := head
	pSecond := head
	//第一个链表先走k-1步
	for k > 1 {
		if pFirst.Next == nil {
			return nil
		}
		pFirst = pFirst.Next
		k--
	}
	//双链表并排走
	for pFirst.Next != nil {
		pFirst = pFirst.Next
		pSecond = pSecond.Next
	}
	return pSecond
}

//ToString 打印单链表
func ToString(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}
