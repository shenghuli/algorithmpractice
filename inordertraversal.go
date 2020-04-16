package main

import (
	"fmt"
)

func main(){
  //二叉树的中序遍历
	root := inordertraversal.TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(InorderTraversalStack(&root))
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//InorderTraversal 中序遍历，递归解法
func InorderTraversal(root *TreeNode) []int {
	retList := []int{}
	if root == nil {
		return retList
	}
	retList = append(retList, InorderTraversal(root.Left)...)
	retList = append(retList, root.Val)
	retList = append(retList, InorderTraversal(root.Right)...)
	return retList
}

//InorderTraversalStack 中序遍历，栈解法
func InorderTraversalStack(root *TreeNode) []int {
	retList := []int{}
	tmpStack := []*TreeNode{}
	for len(tmpStack) > 0 || root != nil {
		if root != nil {
			tmpStack = append(tmpStack, root)
			root = root.Left
		} else {
			stackTop := tmpStack[len(tmpStack)-1]
			tmpStack = tmpStack[:len(tmpStack)-1]
			retList = append(retList, stackTop.Val)
			root = stackTop.Right
		}
	}
	return retList
}

func PrintQueue(q []*TreeNode) {
	for i := 0; i < len(q); i++ {
		if q[i] == nil {
			continue
		}
		fmt.Print(q[i].Val, " ")
	}
	fmt.Println("")
}
