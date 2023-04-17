package _94_traverse_bt_inorder

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TraverseRec(node *TreeNode, traverse *[]int) {
	if node == nil {
		return
	}
	TraverseRec(node.Left, traverse)
	*traverse = append(*traverse, node.Val)
	TraverseRec(node.Right, traverse)

}

func inorderTraversal(root *TreeNode) []int {
	traverse := &[]int{}
	TraverseRec(root, traverse)
	return *traverse
}
