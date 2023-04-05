package utils

type ListNode struct {
	Val  int
	Next *ListNode
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
