package _404_sum_of_left_leaves

import (
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	tt := []struct {
		name    string
		message string
		input   []int
		output  int
	}{
		{
			"base case",
			"There are two left leaves in the binary tree, leafes are nodes WITH NO CHILDREN.",
			[]int{3, 9, 20, NULL, NULL, 15, 7},
			24,
		},
		{
			"case_no_leaves",
			"If the tree consists of only root, there is no leaves",
			[]int{1},
			0,
		},
		{
			"case_broken",
			"leetcode test 29/100",
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			"case_one_chile_nil_other_not",
			"test case to prevent nil pointer dereference",
			[]int{1, 2},
			2,
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			inputTree := BinaryTreeFromSlice(test.input)
			got := sumOfLeftLeaves(inputTree)
			if got != test.output {
				t.Errorf("got: %d, want: %d, inputSlice: %v", got, test.output, test.input)
			}
		})
	}
}

func BinaryTreeFromSlice(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: input[0],
	}

	queue := make([]*TreeNode, 1, len(input))
	queue[0] = root

	for i := 1; i < len(input); {
		node := queue[0]
		queue = queue[1:]

		if i < len(input) && input[i] != NULL {
			node.Left = &TreeNode{Val: input[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(input) && input[i] != NULL {
			node.Right = &TreeNode{Val: input[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
