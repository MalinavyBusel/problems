package _1302_deepest_leaves_sum

import "testing"

func Test_deepestLeavesSum(t *testing.T) {
	tt := []struct {
		name    string
		message string
		input   []int
		output  int
	}{
		{
			"case_2",
			"the function should return the sum of the deepest leaves' values",
			[]int{1, 2, 3, 4, 5, NULL, 6, 7, NULL, NULL, NULL, NULL, 8},
			15,
		},
		{
			"case_1",
			"the function should return the sum of the deepest leaves' values",
			[]int{6, 7, 8, 2, 7, 1, 3, 9, NULL, 1, 4, NULL, NULL, NULL, 5},
			19,
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			inputTree := BinaryTreeFromSlice(test.input)
			got := deepestLeavesSum(inputTree)
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

func (root *TreeNode) ToSlice() []int {
	if root == nil {
		return []int{}
	}

	queue, result := make([]*TreeNode, 1), make([]int, 0, 16)

	queue[0] = root
	l := 1

	for len(queue) != 0 {
		if cur := queue[0]; cur != nil {
			result = append(result, cur.Val)
			l = len(result)
			queue = append(queue, cur.Left, cur.Right)
		} else {
			result = append(result, NULL)
		}
		queue = queue[1:]
	}

	return result[:l]
}
