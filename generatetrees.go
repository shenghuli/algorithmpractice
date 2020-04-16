package main

import (
	"fmt"
)


func main() {

	//不同的二叉搜索树 II
	l := GenerateTrees(3)
	fmt.Println(l)
}

}










//95. 不同的二叉搜索树 II
//给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
//
//示例:
//
//输入: 3
//输出:
//[
//[1,null,3,2],
//[3,2,null,1],
//[3,1,null,null,2],
//[2,1,3],
//[1,null,2,null,3]
//]
//解释:
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//1         3     3      2      1
//\       /     /      / \      \
//3     2     1      1   3      2
///     /       \                 \
//2     1         2                 3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTrees(1, n)
}

func generateTrees(start, end int) []*TreeNode {
	list := []*TreeNode{}
	if start > end {
		list = append(list, nil)
		return list
	}
	if start == end {
		tmpNode := &TreeNode{
			Val: start,
		}
		list = append(list, tmpNode)
		return list
	}
	for i := start; i <= end; i++ {
		leftTrees := generateTrees(start, i-1)
		rightTrees := generateTrees(i+1, end)
		for l := 0; l < len(leftTrees); l++ {
			for r := 0; r < len(rightTrees); r++ {
				root := TreeNode{
					Val:   i,
					Left:  leftTrees[l],
					Right: rightTrees[r],
				}
				list = append(list, &root)
			}
		}
	}
	return list
}
