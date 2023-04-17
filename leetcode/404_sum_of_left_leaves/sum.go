package _404_sum_of_left_leaves

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRecursive(node *TreeNode, isLeft int) int {
	if node == nil || node.Val == NULL {
		return 0
	}
	if (node.Left == nil || node.Left.Val == NULL) && (node.Right == nil || node.Right.Val == NULL) {
		return node.Val * isLeft
	}
	return sumRecursive(node.Left, 1) + sumRecursive(node.Right, 0)
}

func sumOfLeftLeaves(node *TreeNode) int {
	return sumRecursive(node, 0)
}
