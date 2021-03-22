package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (l3 *ListNode) {
	var tmp *ListNode
	n1, n2, carry := 0, 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		var v = n1 + n2 + carry
		if v >= 10 {
			carry = v/10
			v = v%10
		} else {
			carry = 0
		}

		if tmp == nil {
			l3 = &ListNode{Val: v}
			tmp = l3
		} else {
			tmp.Next = &ListNode{Val: v}
			tmp = tmp.Next
		}
	}
	return l3
}
