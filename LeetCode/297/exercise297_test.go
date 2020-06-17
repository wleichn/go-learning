package _97

import (
	"fmt"
	"testing"
)

/**
 * CreatedBy: wlei at 2020/6/16
 * Description:
 *
 */
func TestSerialize(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1

	node2 := new(TreeNode)
	node2.Val = 2
	node2.Right = nil
	node2.Left = nil

	node3 := new(TreeNode)
	node3.Val = 3
	root.Left = node2
	root.Right = node3

	node4 := new(TreeNode)
	node4.Val = 4
	node4.Right = nil
	node4.Left = nil

	node5 := new(TreeNode)
	node5.Val = 5
	node5.Right = nil
	node5.Left = nil

	node3.Left = node4
	node3.Right = node5

	codec := new(Codec)
	res := codec.serialize(root)

	fmt.Print(res)
}
