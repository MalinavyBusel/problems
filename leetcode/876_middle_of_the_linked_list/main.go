package _876_middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	last := &ListNode{Next: head}
	middle := &ListNode{Next: head}
	flag := 0
	for last != nil {
		if flag%2 == 0 {
			middle = middle.Next
		}
		last = last.Next
		flag++
	}
	return middle
}
