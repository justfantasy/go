package main

import (
	"fmt"
	. "p2/leetcode"
)


func main()  {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(AddTwoNumbers(l1, l2))
	//fmt.Println(l.TwoSum([]int{1, 6, 3}, 9))
}
