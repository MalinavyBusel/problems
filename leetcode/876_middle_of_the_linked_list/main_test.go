package _876_middle_of_the_linked_list

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		list    []int
		want    []int
	}{
		{
			name:    "case1",
			message: "the func should return the link to the middle node of the list",
			list:    []int{1, 2, 3, 4, 5},
			want:    []int{3, 4, 5},
		},
		{
			name:    "case_even_nodes",
			message: "if the amount of nodes is even, the func should return the one from the right",
			list:    []int{1, 2, 3, 4, 5, 6},
			want:    []int{4, 5, 6},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := middleNode(LinkedListFromSlice(test.list)).ToSlice()
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf("got: %q, want: %q", actual, test.want)
			}
		})

	}
}

func LinkedListFromSlice(input []int) *ListNode {
	dummy := new(ListNode)

	cur := dummy
	for _, n := range input {
		next := &ListNode{Val: n}
		cur, cur.Next = next, next
	}

	return dummy.Next
}

func (ln *ListNode) ToSlice() []int {
	slice := make([]int, 0)
	for cur := ln; cur != nil; cur = cur.Next {
		slice = append(slice, cur.Val)
	}

	return slice
}
