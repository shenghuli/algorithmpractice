
/*
 题目描述：两数相加
 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

package main

//ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}


//AddTwoNumbers 两数求和
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	Head := &ListNode{}
	carry := 0
	if l1 != nil {
		Head.Val = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		Head.Val += l2.Val
		l2 = l2.Next
	}
	carry = Head.Val / 10
	Head.Val = Head.Val % 10
	NodePtr := Head
	for l1 != nil || l2 != nil {
		tmpNode := ListNode{}
		tmpNode.Val = carry
		if l1 != nil {
			tmpNode.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmpNode.Val += l2.Val
			l2 = l2.Next
		}
		//修正值
		carry = tmpNode.Val / 10
		tmpNode.Val = tmpNode.Val % 10
		NodePtr.Next = &tmpNode
		NodePtr = &tmpNode
	}
	if carry != 0 {
		NodePtr.Next = &ListNode{
			Val: carry,
		}
	}
	return Head
}

//PrintListNode 打印ListNode
func PrintListNode(l *ListNode) {
	lst := []int{}
	for ; l != nil; l = l.Next {
		lst = append(lst, l.Val)
	}
	count := len(lst)
	for i := count - 1; i >= 0; i-- {
		fmt.Print(lst[i])
	}
	fmt.Print("\n")
}


func main() {
	Node3 := ListNode{
		Val: 3,
	}
	Node4 := ListNode{
		Val:  4,
		Next: &Node3,
	}
	l1 := ListNode{
		Val:  2,
		Next: &Node4,
	}

	Node2_4 := ListNode{
		Val: 9,
	}
	Node6 := ListNode{
		Val:  6,
		Next: &Node2_4,
	}
	l2 := ListNode{
		Val:  9,
		Next: &Node6,
	}
	PrintListNode(&l1)
	PrintListNode(&l2)
	tmpNodeList := AddTwoNumbers(&l1, &l2)
	PrintListNode(tmpNodeList)
}





