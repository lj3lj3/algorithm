package jz

// 链表中环的入口结点
//
// 一个链表中包含环，请找出该链表的环的入口结点。要求不能使用额外的空间。
//
// 1. Find count of nodes in loop by using two pointers, p1 forwards one node per step, p2 forwards 2 nodes per step,
// when p1 meets p2, this node should be in the loop, then loop the loop until we found this node again. And the nodes we passed
// is the count of nodes in the loop, save it as n
//
// 2. Reset p1 and p2, let p1 skip n nodes, then move forward with p2, we p1 and p2 meet, that node is the entry of loop
func EntryNodeOfLoop(head *LinkNode) *LinkNode {
	// Find the meeting point
	p1, p2 := head, head.next // let p2 move a step ahead
	for p1.value != p2.value {
		// there is no loop in this linked list
		if p2.next == nil || p2.next.next == nil {
			return nil
		}
		p1 = p1.next
		p2 = p2.next.next
	}

	// Find the count of nodes in the loop
	nodesCount := 1
	p2 = p2.next
	for p1.value != p2.value {
		p2 = p2.next
		nodesCount++
	}

	// Reset pointers
	p1, p2 = head, head
	// Move p2 forward
	for i := 0; i < nodesCount; i++ {
		p2 = p2.next
	}
	// Move together
	for p1.value != p2.value {
		p1 = p1.next
		p2 = p2.next
	}

	// Entry node
	return p1
}
