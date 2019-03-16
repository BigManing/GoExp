package algorithm

import "testing"

func TestDeleteRepeatDataInLink(t *testing.T) {
	header := &LinkNode{Name: "A"}
	header.Next = &LinkNode{Name: "B"}
	header.Next.Next = &LinkNode{Name: "B"}
	header.Next.Next.Next = &LinkNode{Name: "C"}

	//删除 重复  数据
	deleteRepeatDataInLink(header)

	// 打印 数据
	current := header

	for current != nil {
		println(current.Name)
		current = current.Next
	}

	//	 期待 打印  ABC
}
