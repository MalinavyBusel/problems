package _21_merge_sorted_linked_lists

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		list1   []int
		list2   []int
		want    []int
	}{
		{
			name:    "base_case",
			message: "the func shoult merge two linked list in ascending order",
			list1:   []int{1, 2, 4},
			list2:   []int{1, 3, 4},
			want:    []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:    "case_empty_lists",
			message: "when merging two empty lists, an empty list should be returned",
			list1:   []int{},
			list2:   []int{},
			want:    []int{},
		},
		{
			name:    "case_one_of_the_lists_ends_first",
			message: "the remaining values should be added to the tail",
			list1:   []int{1, 2, 4},
			list2:   []int{},
			want:    []int{1, 2, 4},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := mergeTwoLists(LinkedListFromSlice(test.list1), LinkedListFromSlice(test.list2)).ToSlice()
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
