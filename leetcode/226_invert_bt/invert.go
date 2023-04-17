package _226_invert_bt

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertRec(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	InvertRec(node.Left)
	InvertRec(node.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	InvertRec(root)
	return root
}
