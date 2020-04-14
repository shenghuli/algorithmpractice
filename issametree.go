package main

import (
	"fmt"
)


func main() {
  //相同的树
	p := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	q := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(issametree.IsSameTree(&p, &q))
}









//100. 相同的树
//给定两个二叉树，编写一个函数来检验它们是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
//示例 1:
//
//输入:       1         1
/// \       / \
//2   3     2   3
//
//[1,2,3],   [1,2,3]
//
//输出: true

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//IsSameTree 判断两颗二叉树是否相同
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	}
	if q == nil {
		return false
	}
	if p.Val == q.Val {
		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}
	return false
}

