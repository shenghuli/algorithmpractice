package main

import (
	"fmt"
)

func main(){
  //K 个一组翻转链表
	root := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reversekgroup.PrintList(ReverseKGroup2(&root, 3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//25. K 个一组翻转链表
//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
//示例：
//
//给你这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
//说明：
//
//你的算法只能使用常数的额外空间。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

//两种解法，栈解法&递归，参考：https://zhuanlan.zhihu.com/p/64238202

//ReverseKGroup 栈解法
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	pHead := &ListNode{}
	prev := pHead
	stack := []*ListNode{}
	for {
		counter := k
		cur := head
		//入栈
		for counter > 0 && cur != nil {
			stack = append(stack, cur)
			cur = cur.Next
			counter--
		}
		if counter != 0 {
			prev.Next = head
			break
		}

		//出栈
		for counter < k {
			prev.Next = stack[k-counter-1]
			prev = prev.Next
			counter++
		}
		stack = []*ListNode{}
		head = cur
	}
	return pHead.Next
}

//ReverseKGroup 递归解法
func ReverseKGroup2(head *ListNode, k int) *ListNode {
	cur := head
	counter := 0
	for counter < k {
		cur = cur.Next
		counter++
	}

	//翻转
	if counter == k {
		cur := ReverseKGroup(cur, k)
		for counter > 0 {
			tmpNode := head.Next
			head.Next = cur
			cur = head
			head = tmpNode
			counter--
		}
		head = cur
	}
	return head
}

//PrintList 打印链表
func PrintList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	fmt.Println()
}

