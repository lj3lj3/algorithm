package jz

// 链表中倒数第 K 个结点
//
// 输入一个链表，输出该链表中倒数第k个结点。
func FindKthToTail(node *LinkNode, k int) *LinkNode {
	// find the first K-th node
	startNode := node
	for i := 0; i < k-1; i++ { // the first node is current node
		if node == nil {
			break
		}
		// move to the next node
		node = node.next
	}
	if node == nil {
		// not exists
		return nil
	}
	// exists
	endNode := node

	// moving both start and end node, until the end node reach the end of linked list
	for endNode.next != nil {
		endNode = endNode.next
		startNode = startNode.next
	}

	// now, the start node is the K-th node in the linked list
	return startNode
}
