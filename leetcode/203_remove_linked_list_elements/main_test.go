package _203_remove_linked_list_elements

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		gotList []int
		gotVal  int
		want    []int
	}{
		{
			name:    "case1",
			message: "the func should remove all the occurences of value from the slice",
			gotList: []int{1, 2, 6, 3, 4, 5, 6},
			gotVal:  6,
			want:    []int{1, 2, 3, 4, 5},
		},
		{
			name:    "case_empty_slice",
			message: "when given an empty list, the returned list should be empty too",
			gotList: []int{},
			gotVal:  6,
			want:    []int{},
		},
		{
			name:    "case_repeating_value",
			message: "when given a list of one repeating given value, the returned list should be empty",
			gotList: []int{7, 7, 7, 7},
			gotVal:  7,
			want:    []int{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := removeElements(LinkedListFromSlice(test.gotList), test.gotVal).ToSlice()
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf("got: %q, want: %q", actual, test.gotList)
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
