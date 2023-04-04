package _83_remove_duplicates_from_sorted_L_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	first := &ListNode{Val: 101}
	last := first
	for node := head; node != nil; node = node.Next {
		if node.Val != last.Val {
			last.Next = node
			last = node
		}
	}
	last.Next = nil
	return first.Next
}
