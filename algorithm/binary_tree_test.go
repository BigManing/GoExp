package algorithm

import (
	"testing"
)

func TestBinaryPrint(t *testing.T) {
	rootNode := generateNode()

	println("先序打印")
	rootNode.binaryPrint(xian_xu)
	//
	//println("中序打印")
	//rootNode.binaryPrint(zhong_xu)
	//
	//println("后序打印")
	//rootNode.binaryPrint(hou_xu)

}

// 生成二叉树  数据
func generateNode() TreeNode {

	rootNode := TreeNode{name: "A"}

	rootNode.left = &TreeNode{name: "B"}
	rootNode.right = &TreeNode{name: "C"}

	rootNode.left.left = &TreeNode{name: "D"}
	rootNode.left.right = &TreeNode{name: "F"}

	rootNode.left.right.left = &TreeNode{name: "E"}

	rootNode.right.left = &TreeNode{name: "G"}
	rootNode.right.left.right = &TreeNode{name: "H"}

	rootNode.right.right = &TreeNode{name: "I"}
	return rootNode
}
