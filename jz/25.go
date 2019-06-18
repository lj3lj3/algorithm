package jz

// 合并两个排序的链表
//
// 输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
func Merge(node1 *LinkNode, node2 *LinkNode) *LinkNode {
	// Create a new link list
	linkList := &LinkNode{}
	head := linkList
	for node1 != nil && node2 != nil {
		if node1.value < node2.value {
			// Append the node1
			linkList.next = node1
			node1 = node1.next
		} else {
			// Append the node2
			linkList.next = node2
			node2 = node2.next
		}
		// Move
		linkList = linkList.next
	}

	// Check if node 1 or node2 is empty
	if node1 != nil {
		linkList.next = node1
	} else if node2 != nil {
		linkList.next = node2
	}

	return head.next
}
