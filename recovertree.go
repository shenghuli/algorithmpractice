package main

import (
	"fmt"
)

func main(){
  //恢复二叉搜索树
	//root := recovertree.TreeNode{
	//	Val: 3,
	//	Left: &recovertree.TreeNode{
	//		Val: 1,
	//	},
	//	Right: &recovertree.TreeNode{
	//		Val: 4,
	//		Left: &recovertree.TreeNode{
	//			Val: 2,
	//		},
	//	},
	//}
	root := recovertree.TreeNode{
		Val: 1,
		Left: &recovertree.TreeNode{
			Val: 3,
			Right: &recovertree.TreeNode{
				Val: 2,
			},
		},
	}
	recovertree.RecoverTree(&root)
	fmt.Println(root)
}


//99. 恢复二叉搜索树
//二叉搜索树中的两个节点被错误地交换。
//
//请在不改变其结构的情况下，恢复这棵树。
//
//示例 1:
//
//输入: [1,3,null,null,2]
//
//1
///
//3
//\
//2
//
//输出: [3,1,null,null,2]
//
//3
///
//1
//\
//2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//利用中序遍历二叉搜素的有序性，需要记录两个点，第一个比后一个小的当前节点和之后比后一个大的节点的节点中的后一个节点

func RecoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	first, second, prev = innerRecoverTree(root, first, second, prev)
	if first == nil || second == nil {
		return
	}
	first.Val, second.Val = second.Val, first.Val
}

//RecoverTree 恢复二叉树，中序遍历法
func innerRecoverTree(root, f, s, p *TreeNode) (first, second, prev *TreeNode) {
	first, second, prev = f, s, p
	if root == nil {
		return
	}
	first, second, prev = innerRecoverTree(root.Left, first, second, prev)
	//处理当前节点,逆序节点出现在左孩子中
	if prev != nil && prev.Val > root.Val {
		if first == nil {
			first = prev
			second = root
		} else {
			second = root
		}
	} else {
		prev = root
	}
	first, second, prev = innerRecoverTree(root.Right, first, second, prev)
	return
}
