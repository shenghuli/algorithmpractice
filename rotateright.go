package main

import (
	"fmt"
)

func main(){
  //旋转链表
	root := rotateright.ListNode{
		Val: 1,
		Next: &rotateright.ListNode{
			Val: 2,
			Next: &rotateright.ListNode{
				Val: 3,
				Next: &rotateright.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(RotateRight(&root, 3))
}

//61. 旋转链表
//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL, k = 2
//输出: 4->5->1->2->3->NULL
//解释:
//向右旋转 1 步: 5->1->2->3->4->NULL
//向右旋转 2 步: 4->5->1->2->3->NULL
//参考：https://zhuanlan.zhihu.com/p/133272673

type ListNode struct {
	Val  int
	Next *ListNode
}

//RotateRight
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pTmp := head
	count := 1
	for pTmp.Next != nil {
		pTmp = pTmp.Next
		count++
	}
	pTmp.Next = head //成环
	k = count - (k % count)

	p1 := head
	p2 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
		p2 = p2.Next
	}

	for i := 1; i < count; i++ {
		p2 = p2.Next
	}
	p2.Next = nil
	return p1
}
