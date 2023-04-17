package _104_max_depth_of_bt

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepthRec(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return MaxInt(MaxDepthRec(node.Left), MaxDepthRec(node.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return MaxDepthRec(root)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
