package algorithm

import "fmt"

// 二叉树 节点
type TreeNode struct {
	name        string
	left, right *TreeNode
}

// 打印类型
type PrintType int

const (
	//先序 打印   从顶点 node开始 打印
	// 先打印 最左侧节点   到底后   再打印最右边的
	//按照根节点->左子树->右子树的顺序访问二叉树
	xian_xu PrintType = iota
	// 中序打印   按照左子树->根节点->右子树的顺序访问
	zhong_xu
	//后序遍历  （1）采用后序递归遍历左子树；（2）采用后序递归遍历右子树；（3）访问根节点
	hou_xu
)

/*
二叉树  遍历 打印
*/
func (node *TreeNode) binaryPrint(pType PrintType) {
	if node == nil {
		return
	}
	if pType == xian_xu {
		fmt.Print(node.name)
	}

	node.left.binaryPrint(pType)

	if pType == zhong_xu {
		fmt.Print(node.name)
	}

	node.right.binaryPrint(pType)

	if pType == hou_xu {
		fmt.Print(node.name)
	}
}
