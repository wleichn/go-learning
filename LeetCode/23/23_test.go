package _3

import (
	"fmt"
	"testing"
)

/**
 * CreatedBy: wlei at 2020/6/9
 * Description:
 *
 */

func TestMergeKLists(t *testing.T) {
	samples := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	results := []int{
		1, 1, 2, 3, 4, 4, 5, 6,
	}

	samplesList := make([]*ListNode, len(samples))
	for index, sample := range samples {
		nodeHead := new(ListNode)
		listPos := nodeHead
		for _, val := range sample {
			tempNode := new(ListNode)
			tempNode.Val = val
			listPos.Next = tempNode
			listPos = tempNode
		}
		nodeHead = nodeHead.Next
		samplesList[index] = nodeHead
	}

	resultHead := new(ListNode)
	resultPos := resultHead
	for _, result := range results {
		tempNode := new(ListNode)
		tempNode.Val = result
		resultPos.Next = tempNode
		resultPos = tempNode
	}
	resultHead = resultHead.Next

	for index, sample := range samplesList {
		fmt.Printf("List %d: %d ", index, sample.Val)
		printList(sample, "sample")
	}
	printList(resultHead, "result")

	resultActual := mergeKLists(samplesList)
	printList(resultActual, "actual result")

}

func printList(resultHead *ListNode, str string) {
	fmt.Printf("list of %s: %d ", str, resultHead.Val)
	for resultHead.Next != nil {
		resultHead = resultHead.Next
		fmt.Printf(" %d ", resultHead.Val)
	}
	fmt.Printf("\n")
}
