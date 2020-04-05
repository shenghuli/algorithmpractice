package main

func main() {
  //删除排序链表中的重复元素
	l1 := &deleteduplicates.ListNode{
		Val: 1,
		Next: &deleteduplicates.ListNode{
			Val: 1,
			Next: &deleteduplicates.ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	ToString(DeleteDuplicates(l1))
}


//ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2

//DeleteDuplicates 删除排序列表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := head.Next
	second := head
	for first != nil {
		if first.Val == second.Val {
			second.Next = first.Next
		} else {
			second = second.Next
		}
		first = first.Next
	}
	return head
}

//DeleteDuplicates 递归解法
func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = DeleteDuplicates2(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}

//ToString 打印单链表
func ToString(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}
