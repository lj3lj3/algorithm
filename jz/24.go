package jz

// 反转链表
//
// 输入一个链表，反转链表后，输出新链表的表头。
//
// 构建一个有一个头结点的列表 然后依次插入节点到头节点的下一个节点
func ReverseList(node *LinkNode) *LinkNode {
	// create a link list with only one node, and inject the node
	head := &LinkNode{}
	for node != nil {
		// append node
		newListNext := head.next
		head.next = node
		nodeNext := node.next   // temporary save next node in the old list
		node.next = newListNext // set next node in the new link list
		node = nodeNext         // set next node back
	}

	return head.next
}
