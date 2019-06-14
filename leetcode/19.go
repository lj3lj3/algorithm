package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Remove Nth head From End of List
//
// Add a start to the head
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{ // append a starter to the head
		Val:  -1,
		Next: head,
	}
	// find the first n+1 th node
	startNode, node := head, head
	for i := 0; i < n; i++ { // we want find the node before n, the first node is the current node
		if node == nil {
			break
		}
		// move to the next head
		node = node.Next
	}
	if node == nil {
		// not exists
		return nil
	}
	// exists
	endNode := node

	// moving both start and end node, until the end node reach the end of linked list
	for endNode.Next != nil {
		endNode = endNode.Next
		startNode = startNode.Next
	}

	// now, the start node is the node a node before n-th node in the linked list
	startNode.Next = startNode.Next.Next

	return head.Next
}
