package algorithm

import "testing"

func TestDeleteRepeatDataInLink(t *testing.T) {
	header := getLinkNode()

	//删除 重复  数据
	deleteRepeatDataInLink(header)

	// 打印 数据

	printLink(header)

	//	 期待 打印  ABC
}

func TestDeleteLink(t *testing.T) {
	node := getLinkNode()
	DeleteLink(&node)
	printLink(node)

}
func getLinkNode() *LinkNode {
	header := &LinkNode{Name: "A"}
	header.Next = &LinkNode{Name: "B"}
	header.Next.Next = &LinkNode{Name: "B"}
	header.Next.Next.Next = &LinkNode{Name: "C"}
	return header
}
func printLink(header *LinkNode) {
	node := header
	for node != nil {
		println(node.Name)
		node = node.Next
	}
}
