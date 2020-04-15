package main

import "fmt"

func main() {

  //对称二叉树
	p := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(issymmetric.IsSymmetricQueue(&p))

}


//对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//IsSymmetric 对称二叉树，递归解法,剑指offer28题解法
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

//IsSymmetricQueue 对称二叉树，队列解法
func IsSymmetricQueue(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//队列
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		node1 := queue[0]
		node2 := queue[1]
		//出队
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left, node2.Right, node1.Right, node2.Left)
	}
	return true
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
