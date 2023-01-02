//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 2828 👎 0

package cn

import (
	. "myself/leetcode/leetcode/editor/cn/utils"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := NewLinkNodeList([]int{1, 2, 4})
	l2 := NewLinkNodeList([]int{1, 3})
	l1Tail := GetLinkListTail(l1)
	l2Tail := GetLinkListTail(l2)
	l2Tail.Next = l1Tail

	header := mergeTwoLists(l1, l2)
	PrintLinkList(header)
}

//leetcode submit region begin(Prohibit modification and deletion)

//Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	var p, p1, p2 *ListNode
	p = dummy
	p1 = l1
	p2 = l2
	for true {
		if p1 == nil {
			p.Next = p2
			return dummy.Next
		}
		if p2 == nil {
			p.Next = p1
			return dummy.Next
		}

		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}

		p = p.Next
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
