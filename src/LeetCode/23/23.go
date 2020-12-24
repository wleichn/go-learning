package _3

import "math"

/**
 * CreatedBy: wlei at 2020/6/9
 * Description:
 *
 * 注意：getMiniNode函数中的传入参数可以不使用 *[]*ListNode这种切片指针类型，因为传入的切片虽然是副本，
 * 但是和函数外部的切片指向同一个数组，所以在函数内部改变切片时可以影响到函数外部的切片
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

	return resultListNode.Next
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
		if node != nil && node.Val < miniVal {
			miniVal = node.Val
			miniPos = index
		}
	}
	if miniPos != -1 {
		miniNode := (lists)[miniPos]
		miniNode = miniNode.Next
		(lists)[miniPos] = miniNode
		return miniVal, true
	}
	return 0, false
}
