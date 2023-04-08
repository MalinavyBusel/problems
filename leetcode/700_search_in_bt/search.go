package _700_search_in_bt

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	node := root
	for {
		if node == nil || node.Val == val {
			return node
		}
		if node.Val < val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
}
