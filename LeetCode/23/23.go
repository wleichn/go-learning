package _3

import "math"

/**
 * CreatedBy: wlei at 2020/6/9
 * Description:
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	resultListNode := new(ListNode)
	position := resultListNode
	for !isEmpty(lists) {
		tempVal, exist := getMiniNode(lists)
		if exist {
			tempNode := new(ListNode)
			tempNode.Val = tempVal
			position.Next = tempNode
			position = position.Next
		}
	}

	return resultListNode
}

func isEmpty(lists []*ListNode) bool {
	for _, node := range lists {
		if node != nil {
			return false
		}
	}

	return true
}

func getMiniNode(lists []*ListNode) (int, bool) {
	miniPos := -1
	miniVal := math.MaxInt32
	for index, node := range lists {
		if node.Val < miniVal {
			miniVal = node.Val
			miniPos = index
		}
	}
	if miniPos != -1 {
		miniNode := lists[miniPos]
		miniNode = miniNode.Next
		return miniVal, true
	}
	return 0, false
}
