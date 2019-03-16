package algorithm

// 链表  结构体
type LinkNode struct {
	Name string
	Next *LinkNode
}

//记录 当前node   判断 node.next  的值
func deleteRepeatDataInLink(node *LinkNode) {
	if node == nil {
		return
	}
	// 记录当前node
	current := node
	for current.Next != nil {
		//判断前后 node  是否为重复的值
		if current.Name == current.Next.Name {
			// 重新 指向 下下个
			current.Next = current.Next.Next
		} else {
			// 指向下一个
			current = current.Next
		}
	}
	println(node)
}
