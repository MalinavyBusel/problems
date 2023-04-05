package _83_remove_duplicates_from_sorted_L_list

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
			message: "the func should return the list where each number occurs not more than once",
			list:    []int{1, 1, 2},
			want:    []int{1, 2},
		},
		{
			name:    "case_empty_lists",
			message: "the func should return the list where each number occurs not more than once",
			list:    []int{1, 1, 2, 3, 3},
			want:    []int{1, 2, 3},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := deleteDuplicates(LinkedListFromSlice(test.list)).ToSlice()
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
