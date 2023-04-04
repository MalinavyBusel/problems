package _1290_binary_to_int

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	numS := ""
	for node := head; node != nil; node = node.Next {
		numS += strconv.Itoa(node.Val)
	}
	val, _ := strconv.ParseInt(numS, 2, 64)
	return int(val)
}
