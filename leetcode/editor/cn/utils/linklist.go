package utils

import (
	"fmt"
)

func NewLinkNode(val int) *ListNode {
	l := &ListNode{
		Val:  val,
		Next: nil,
	}
	return l
}

func NewLinkNodeList(arr []int) *ListNode {
	header := &ListNode{
		Val:  0,
		Next: nil,
	}

	var p = header
	for _, val := range arr {
		l := &ListNode{
			Val:  val,
			Next: nil,
		}
		p.Next = l
		p = p.Next
	}

	return header.Next
}

func GetLinkListTail(header *ListNode) *ListNode {
	p := header
	for true {
		if p.Next == nil {
			return p
		}
		p = p.Next
	}
	return nil
}

func GetLinkListNodeByIndex(header *ListNode, index int) *ListNode {
	p := header
	for i := 0; i < index; i++ {
		if p.Next == nil {
			return nil
		}
		p = p.Next
	}
	return p
}

func PrintLinkList(header *ListNode) {
	var p = header
	res := make([]int, 0)
	for true {
		if p == nil {
			fmt.Println(res)
			break
		}

		res = append(res, p.Val)
		p = p.Next
	}
	return
}
