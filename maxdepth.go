package main

import (
	"fmt"
)

func main() {
  //二叉树的最大深度
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(MaxDepthQueue(root))
}










type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//MaxDepth 二叉树最大深度，剑指offer面试题55解法
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	llen := MaxDepth(root.Left)
	rlen := MaxDepth(root.Right)
	if llen < rlen {
		return rlen + 1
	}
	return llen + 1
}

//MaxDepthQueue 队列解法
func MaxDepthQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue) //本层节点数
		count := 0
		for ; count < size; count++ {
			head := queue[count]
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		//删除本层节点
		queue = queue[count:]
	}
	return depth
}

//PrintQueue 打印队列
func PrintQueue(q []*TreeNode) {
	for i := 0; i < len(q); i++ {
		if q[i] == nil {
			continue
		}
		fmt.Print(q[i].Val, " ")
	}
	fmt.Println("")
}
