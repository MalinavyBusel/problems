package _1302_deepest_leaves_sum

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxIndex(a, b int) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	}
	return 0
}

func deepestLeavesSumRec(node *TreeNode) [2]int {
	if node == nil {
		return [2]int{math.MinInt64, 0}
	}
	if node.Left == nil && node.Right == nil {
		return [2]int{0, node.Val}
	}
	a := deepestLeavesSumRec(node.Left)
	b := deepestLeavesSumRec(node.Right)
	if MaxIndex(a[0], b[0]) > 0 {
		a[0]++
		return a
	} else if MaxIndex(a[0], b[0]) < 0 {
		b[0]++
		return b
	} else {
		a[0]++
		a[1] += b[1]
		return a
	}
}

func deepestLeavesSum(root *TreeNode) int {
	sumArr := deepestLeavesSumRec(root)
	return sumArr[1]
}
