package main

import (
	"fmt"
)

func main() {
	//环形列表2
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l1,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l1.Next = l2
	fmt.Println(DetectCycle2(l3))
}



//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//示例 1：
//
//输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。

//ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//DetectCycle hash解法
func DetectCycle(head *ListNode) *ListNode {
	NodeMap := map[*ListNode]bool{}
	for head != nil {
		if _, ok := NodeMap[head]; ok {
			return head
		}
		NodeMap[head] = true
		head = head.Next
	}

	return nil
}

//DetectCycle2 剑指offer第23题解法
func DetectCycle2(head *ListNode) *ListNode {
	pFast := head
	pSlow := head
	var meetingNode *ListNode = nil
	for pFast != nil && pFast.Next != nil {
		pFast = pFast.Next.Next
		pSlow = pSlow.Next
		if pFast == pSlow { //找到了
			meetingNode = pFast
			break
		}
	}
	//没环
	if meetingNode == nil {
		return nil
	}

	//计算环的长度
	pTmpNode := meetingNode
	nodeListLen := 1
	for pTmpNode.Next != meetingNode {
		nodeListLen++
		pTmpNode = pTmpNode.Next
	}
	//双指针
	pFirst := head
	//第一个指针先走环长度步,两个人的距离是一个周长，而第二个人一步一步的走，一定能走到
	//环的入口点，此时第一个人刚好比第二个人多走了一圈，同时出现环的入口点
	for i := 0; i < nodeListLen; i++ {
		pFirst = pFirst.Next
	}
	pSecond := head
	for pFirst != pSecond {
		pFirst = pFirst.Next
		pSecond = pSecond.Next
	}
	return pFirst
}
