package _97

/**
 * CreatedBy: wlei at 2020/6/16
 * Description:
 *
 */

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return *new(Codec)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res string
	p := root
	q := new(Queue)
	q.EnQueue(p)

	for !q.IsEmpty() {
		node := q.DeQueue()
		switch t := node.(type) {
		case nil:
			res += " "
		case *TreeNode:
			res += strconv.Itoa(t.Val)
			q.EnQueue(t.Left)
			q.EnQueue(t.Right)
		default:
			break
		}
	}
	fmt.Print(res)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var root *TreeNode
	if len(data) == 0 || data[0] == ' ' {
		return root
	}
	node := new(TreeNode)
	root = node
	node.Val = int(data[0])
	queue := new(Queue)
	queue.EnQueue(node)
	for index := 1; index < len(data); {
		if !queue.IsEmpty() {
			generics := queue.DeQueue()
			switch t := generics.(type) {
			case *TreeNode:
				if data[index] == ' ' {
					t.Left = nil
				} else {
					t.Left = new(TreeNode)
					t.Val = int(data[index])
					queue.EnQueue(t.Left)
				}
				index++
				if index >= len(data) {
					break
				}

				if data[index] == ' ' {
					t.Right = nil
				} else {
					t.Right = new(TreeNode)
					t.Val = int(data[index])
					queue.EnQueue(t.Right)
				}
				index++
			}
		}
	}
	return root
}

func PreorderErgodic(root *TreeNode) (res string) {
	p := root
	if p == nil {
		return ""
	}

	res += string(p.Val)
	if p.Left != nil {
		res += PreorderErgodic(p.Left)
	}
	if p.Right != nil {
		res += PreorderErgodic(p.Right)
	}
	return
}

type Generics interface{}
type Queue struct {
	data []Generics
	len  int
}

func (q *Queue) EnQueue(element Generics) {
	if q == nil {
		return
	}
	q.data = append(q.data, element)
	q.len++
}
func (q *Queue) DeQueue() (res Generics) {
	if q == nil || q.len == 0 {
		return
	}

	res = q.data[0]
	q.data = q.data[1:]
	q.len--
	return
}
func (q *Queue) IsEmpty() bool {
	if q == nil || q.len == 0 {
		return true
	}
	return false
}
func (q *Queue) Size() int {
	if q == nil {
		return -1
	}
	return q.len
}
