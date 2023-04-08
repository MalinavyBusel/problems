package _94_traverse_bt_inorder

import (
	"reflect"
	"testing"
)

func Test_TraverseInorder(t *testing.T) {
	tt := []struct {
		name    string
		message string
		input   []int
		output  []int
	}{
		{
			"base_case_1",
			"the function should invert all children of a node",
			[]int{1, NULL, 2, 3},
			[]int{1, 3, 2},
		},
		{
			"case_one_node",
			"traversing one elem is this elem",
			[]int{1},
			[]int{1},
		},
		{
			"case_empty_tree",
			"traversing empty tree is still an empty tree",
			[]int{},
			[]int{},
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			inputTree := BinaryTreeFromSlice(test.input)
			got := inorderTraversal(inputTree)
			if !reflect.DeepEqual(got, test.output) {
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
