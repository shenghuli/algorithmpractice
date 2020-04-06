package main

import (
  "fmt"
)


func main (){
  //删除链表的倒数第N个节点
	l1 := &removenthfromend.ListNode{
		Val: 1,
		Next: &removenthfromend.ListNode{
			Val: 1,
			Next: &removenthfromend.ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	ToString(RemoveNthFromEnd(l1, 2))
}


//19. 删除链表的倒数第N个节点
//
//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.

//ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//RemoveNthFromEnd 删除链表的倒数第N个节点 求链表长度解法
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//获取链表的长度
	listLen := 0
	tmpNodeList := head
	for tmpNodeList != nil {
		listLen++
		tmpNodeList = tmpNodeList.Next
	}
	index := listLen - n
	//处理删除的是列表头节点的case
	if index == 0 {
		head = head.Next
		return head
	}

	//正常case，寻找要删除链表节点的上一个节点
	tmpNodeList = head
	for i := 0; i < index-1; i++ {
		tmpNodeList = tmpNodeList.Next
	}
	tmpNodeList.Next = tmpNodeList.Next.Next
	return head
}

//RemoveNthFromEnd 双指针解法
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	pFirst := head
	pSecond := head
	//第一个指针先走n步,走到要删除节点的前一个节点
	for ; n > 0; n-- {
		//防止n非法
		if pFirst == nil {
			return nil
		}
		pFirst = pFirst.Next
	}
	//删除链表头节点
	if pFirst == nil {
		head = head.Next
		return head
	}
	//双指针并排走,终止条件为第一个指针走到链表尾
	for pFirst.Next != nil { //此处判断next不为nil，终止时，第二个链表走到要删除节点的前一个节点
		pFirst = pFirst.Next
		pSecond = pSecond.Next
	}
	//删除该节点
	pSecond.Next = pSecond.Next.Next

	return head
}

//ToString 打印单链表
func ToString(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}
